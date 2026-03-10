package pack

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/k9xR7vA2/recharge-common/plan/types"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/crypto/pbkdf2"
)

const viPacksAPIURL = "https://www.myvi.in/wildfly/consumer/api/vodafoneidea/web"

// VIPacksResponse VI 套餐响应
type VIPacksResponse struct {
	STATUS     string      `json:"STATUS"`
	StatusCode int         `json:"statusCode"`
	Data       VIPacksData `json:"data"`
}

// VIPacksData VI 套餐数据
type VIPacksData struct {
	Status              string   `json:"status"`
	Brand               string   `json:"brand"`
	CustomerStatus      string   `json:"customerStatus"`
	SubscriberType      string   `json:"subscriberType"`
	Circle              string   `json:"circle"`
	CircleId            string   `json:"circleId"`
	RecommendedPackList []VIPack `json:"recommendedPackList"`
}

// VIPack 单个套餐
type VIPack struct {
	ProductCategory     string `json:"PRODUCT-CATEGORY"`
	CatTab              string `json:"catTab"`
	UnitCost            string `json:"UNIT_COST"`
	Description         string `json:"description"`
	ItemID              string `json:"ITEM_ID"`
	PackIdentifier      string `json:"packidentifier"`
	WebValidity         string `json:"WEB-VALIDITY"`
	WebSellable         string `json:"WEB-SELLABLE"`
	ReadMore            string `json:"READ_MORE"`
	RechargeSubtype     string `json:"RECHARGE_SUBTYPE"`
	EtopupTypeAttr      string `json:"ETOPUPTYPE_ATTR"`
	OpenMarketFlag      string `json:"OPEN_MARKET_FLAG"`
	DisplayRank         string `json:"DISPLAY-RANK"`
	PromotionTitle      string `json:"PROMOTION_TITLE"`
	UupCircle           string `json:"UUP_CIRCLE"`
	ProductName         string `json:"PRODUCT-NAME"`
	IsHarmonized        string `json:"IS_HARMONIZED"`
	VoiceLine1          string `json:"VOICE_LINE_1"`
	VoiceLine2          string `json:"VOICE_LINE_2"`
	ServiceValidityAttr string `json:"SERVICEVALIDITY_ATTR"`
	DataLine1           string `json:"DATA_LINE_1"`
	ServiceValidity     string `json:"SERVICE_VALIDITY"`
	ValidityAttr        string `json:"VALIDITY_ATTR"`
	SMSLine1            string `json:"SMS_LINE_1"`
}

// viEncryptResult 加密结果
type viEncryptResult struct {
	Params string
	Sl     string
	Algf   string
	Sps    string
}

// VIPackService VI 套餐服务
type VIPackService struct {
	*BasePackService
	client *resty.Client
}

// NewVIPackService 创建 VI 套餐服务
func NewVIPackService() *VIPackService {
	jar, _ := cookiejar.New(nil)
	client := resty.New().
		SetCookieJar(jar).
		SetTimeout(30*time.Second).
		SetHeader("Accept", "application/json, text/plain, */*").
		SetHeader("Origin", "https://www.myvi.in").
		SetHeader("Referer", "https://www.myvi.in/prepaid/online-mobile-recharge").
		SetHeader("methodName", "numberValidation").
		SetHeader("X-Requested-With", "XMLHttpRequest"). // ← 补上
		SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/144.0.0.0 Safari/537.36").
		SetHeader("Sec-Fetch-Dest", "empty").
		SetHeader("Sec-Fetch-Mode", "cors").
		SetHeader("Sec-Fetch-Site", "same-origin")

	return &VIPackService{
		BasePackService: NewBasePackService(types.SourceVI),
		client:          client,
	}
}
func (s *VIPackService) validateAndSeedCookie(ctx context.Context, phoneNumber string) error {
	enc, err := viEncryptMobile(phoneNumber)
	if err != nil {
		return err
	}
	bodyMap := map[string]string{
		"params": enc.Params,
		"sl":     enc.Sl,
		"algf":   enc.Algf,
		"sps":    enc.Sps,
	}
	bodyJSON, _ := json.Marshal(bodyMap)

	resp, err := s.client.R().
		SetContext(ctx).
		SetMultipartFormData(map[string]string{ // ← 改为 multipart
			"mobile": string(bodyJSON),
		}).
		SetHeader("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
		Post("https://www.myvi.in/bin/selected/dxlnumbervalidation")

	if err != nil {
		return err
	}
	fmt.Printf("[VI] validate 响应状态: %d\n", resp.StatusCode())
	fmt.Printf("[VI] validate 响应体: %s\n", string(resp.Body()))
	fmt.Printf("[VI] validate Set-Cookie: %v\n", resp.Header().Values("Set-Cookie"))

	if resp.StatusCode() != 200 {
		return fmt.Errorf("validate status: %d", resp.StatusCode())
	}
	return nil
}

// GetPacks 获取套餐列表
func (s *VIPackService) GetPacks(ctx context.Context, phoneNumber string) ([]types.UnifiedPack, error) {
	// 1. 先调验证接口，让服务端种下 storeCookie / lobCookie
	if err := s.validateAndSeedCookie(ctx, phoneNumber); err != nil {
		return nil, fmt.Errorf("seed cookie failed: %w", err)
	}
	// 1. 加密手机号
	enc, err := viEncryptMobile(phoneNumber)
	if err != nil {
		return nil, fmt.Errorf("encrypt failed: %w", err)
	}

	// 2. 构造 multipart body
	bodyMap := map[string]string{
		"params": enc.Params,
		"sl":     enc.Sl,
		"algf":   enc.Algf,
		"sps":    enc.Sps,
	}
	bodyJSON, _ := json.Marshal(bodyMap)

	fmt.Printf("[VI] 请求套餐 URL: %s\n", viPacksAPIURL)
	fmt.Printf("[VI] 请求 body: mobile=%s\n", string(bodyJSON))
	// 3. 发起请求（multipart/form-data）
	resp, err := s.client.R().
		SetContext(ctx).
		SetMultipartFormData(map[string]string{
			"mobile": string(bodyJSON),
		}).
		Post(viPacksAPIURL)

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	fmt.Printf("[VI] 响应状态: %d\n", resp.StatusCode())
	fmt.Printf("[VI] 响应头: %v\n", resp.Header())
	fmt.Printf("[VI] 响应体: %s\n", string(resp.Body()))
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode())
	}

	// 4. 解析响应
	var viResp VIPacksResponse
	if err := json.Unmarshal(resp.Body(), &viResp); err != nil {
		return nil, fmt.Errorf("parse response failed: %w, body: %s", err, string(resp.Body()))
	}
	if viResp.STATUS != "SUCCESS" {
		return nil, fmt.Errorf("api returned status: %s", viResp.STATUS)
	}

	return s.parsePacks(viResp.Data.RecommendedPackList), nil
}

// CheckAmountExists 检查金额是否存在
func (s *VIPackService) CheckAmountExists(ctx context.Context, phoneNumber, amount string) (bool, []types.UnifiedPack, error) {
	packs, err := s.GetPacks(ctx, phoneNumber)
	if err != nil {
		return false, nil, err
	}
	exists, matched := s.CheckAmountExistsFromPacks(packs, amount)
	return exists, matched, nil
}

// GetPackByAmount 根据金额获取套餐
func (s *VIPackService) GetPackByAmount(ctx context.Context, phoneNumber, amount string) (*types.UnifiedPack, error) {
	packs, err := s.GetPacks(ctx, phoneNumber)
	if err != nil {
		return nil, err
	}
	return s.GetPackByAmountFromPacks(packs, amount), nil
}

// parsePacks 转换为统一套餐格式
func (s *VIPackService) parsePacks(viPacks []VIPack) []types.UnifiedPack {
	seen := make(map[string]bool)
	var packs []types.UnifiedPack

	for _, p := range viPacks {
		if p.WebSellable != "1" {
			continue
		}
		if seen[p.ItemID] {
			continue
		}
		seen[p.ItemID] = true

		packs = append(packs, types.UnifiedPack{
			PackID:      p.ItemID,
			Amount:      p.UnitCost,
			Description: p.ProductName,
			Validity:    p.WebValidity,
			Category:    p.ProductCategory,
			Type:        p.RechargeSubtype,
			Data:        p.DataLine1,
			CarrierCode: s.GetPackSource(),
			RawData:     p,
		})
	}
	return packs
}

// ── 加密逻辑（与 vi.go 保持一致）──

func viEncryptMobile(mobile string) (*viEncryptResult, error) {
	salt := make([]byte, 16)
	iv := make([]byte, 16)
	passPhrase := make([]byte, 16)
	for _, b := range [][]byte{salt, iv, passPhrase} {
		if _, err := rand.Read(b); err != nil {
			return nil, fmt.Errorf("rand failed: %w", err)
		}
	}

	passPhraseHex := hex.EncodeToString(passPhrase)
	key := pbkdf2.Key([]byte(passPhraseHex), salt, 100, 16, sha1.New)

	plaintext, err := json.Marshal(map[string]string{"mobNumber": mobile})
	if err != nil {
		return nil, err
	}

	encrypted, err := viAesCBCEncrypt(key, iv, plaintext)
	if err != nil {
		return nil, err
	}

	return &viEncryptResult{
		Params: url.QueryEscape(encrypted),
		Sl:     hex.EncodeToString(salt),
		Algf:   hex.EncodeToString(iv),
		Sps:    hex.EncodeToString(passPhrase),
	}, nil
}

func viAesCBCEncrypt(key, iv, plaintext []byte) (string, error) {
	blockSize := aes.BlockSize
	padding := blockSize - len(plaintext)%blockSize
	padded := make([]byte, len(plaintext)+padding)
	copy(padded, plaintext)
	for i := len(plaintext); i < len(padded); i++ {
		padded[i] = byte(padding)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, len(padded))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(ciphertext, padded)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
