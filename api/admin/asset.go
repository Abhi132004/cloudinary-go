package admin

// Enables you to manage the assets in your account or cloud.
//
// https://cloudinary.com/documentation/admin_api#resources

import (
	"context"
	"github.com/cloudinary/cloudinary-go/api"
	"time"
)

type AssetParams struct {
	AssetType             api.AssetType    `json:"-"`
	DeliveryType          api.DeliveryType `json:"-"`
	PublicID              string           `json:"-"`
	Exif                  bool             `json:"exif,omitempty"`
	Colors                bool             `json:"colors,omitempty"`
	Faces                 bool             `json:"faces,omitempty"`
	QualityAnalysis       bool             `json:"quality_analysis,omitempty"`
	ImageMetadata         bool             `json:"image_metadata,omitempty"`
	Phash                 bool             `json:"phash,omitempty"`
	Pages                 bool             `json:"pages,omitempty"`
	AccessibilityAnalysis bool             `json:"accessibility_analysis,omitempty"`
	CinemagraphAnalysis   bool             `json:"cinemagraph_analysis,omitempty"`
	Coordinates           bool             `json:"coordinates,omitempty"`
	MaxResults            int              `json:"max_results,omitempty"`
	DerivedNextCursor     string           `json:"derived_next_cursor,omitempty"`
}

//Returns the details of the specified asset and all its derived resources.
//
//Note that if you only need details about the original resource,
//you can also use the uploader.Upload or uploader.Explicit methods, which return the same information and
//are not rate limited.
//
//https://cloudinary.com/documentation/admin_api#get_the_details_of_a_single_resource
func (a *Api) Asset(ctx context.Context, params AssetParams) (*AssetResult, error) {
	res := &AssetResult{}
	_, err := a.get(ctx, api.BuildPath(Assets, params.AssetType, params.DeliveryType,
		params.PublicID), params, res)

	return res, err
}

type AssetResult struct {
	AssetID               string                      `json:"asset_id"`
	PublicID              string                      `json:"public_id"`
	Format                string                      `json:"format"`
	Version               int                         `json:"version"`
	ResourceType          string                      `json:"resource_type"`
	Type                  string                      `json:"type"`
	CreatedAt             time.Time                   `json:"created_at"`
	Bytes                 int                         `json:"bytes"`
	Width                 int                         `json:"width"`
	Height                int                         `json:"height"`
	Backup                bool                        `json:"backup"`
	AccessMode            string                      `json:"access_mode"`
	URL                   string                      `json:"url"`
	SecureURL             string                      `json:"secure_url"`
	Metadata              api.Metadata                `json:"metadata,omitempty"`
	Tags                  []string                    `json:"tags"`
	NextCursor            string                      `json:"next_cursor"`
	Derived               []interface{}               `json:"derived"`
	Etag                  string                      `json:"etag"`
	ImageMetadata         ImageMetadataResult         `json:"image_metadata"`
	Coordinates           struct{}                    `json:"coordinates"`
	Exif                  struct{}                    `json:"exif"`
	Faces                 [][]int                     `json:"faces"`
	IllustrationScore     float64                     `json:"illustration_score"`
	SemiTransparent       bool                        `json:"semi_transparent"`
	Grayscale             bool                        `json:"grayscale"`
	Colors                [][]interface{}             `json:"colors"`
	Predominant           PredominantResult           `json:"predominant"`
	Phash                 string                      `json:"phash"`
	QualityAnalysis       QualityAnalysisResult       `json:"quality_analysis"`
	QualityScore          float64                     `json:"quality_score"`
	AccessibilityAnalysis AccessibilityAnalysisResult `json:"accessibility_analysis"`
	Pages                 int                         `json:"pages"`
	CinemagraphAnalysis   CinemagraphAnalysisResult   `json:"cinemagraph_analysis"`
	Usage                 struct{}                    `json:"usage"`
	OriginalFilename      string                      `json:"original_filename"`
	Error                 api.ErrorResp               `json:"error,omitempty"`
}

type QualityAnalysisResult struct {
	JpegQuality       float64 `json:"jpeg_quality"`
	JpegChroma        float64 `json:"jpeg_chroma"`
	Focus             float64 `json:"focus"`
	Noise             float64 `json:"noise"`
	Contrast          float64 `json:"contrast"`
	Exposure          float64 `json:"exposure"`
	Saturation        float64 `json:"saturation"`
	Lighting          float64 `json:"lighting"`
	PixelScore        float64 `json:"pixel_score"`
	ColorScore        float64 `json:"color_score"`
	Dct               float64 `json:"dct"`
	Blockiness        float64 `json:"blockiness"`
	ChromaSubsampling float64 `json:"chroma_subsampling"`
	Resolution        float64 `json:"resolution"`
}

type AccessibilityAnalysisResult struct {
	ColorblindAccessibilityAnalysis struct {
		DistinctEdges      float64  `json:"distinct_edges"`
		DistinctColors     float64  `json:"distinct_colors"`
		MostIndistinctPair []string `json:"most_indistinct_pair"`
	} `json:"colorblind_accessibility_analysis"`
	ColorblindAccessibilityScore float64 `json:"colorblind_accessibility_score"`
}

type CinemagraphAnalysisResult struct {
	CinemagraphScore float64 `json:"cinemagraph_score"`
}

type ImageMetadataResult map[string]string

type PredominantResult struct {
	Google     [][]interface{} `json:"google"`
	Cloudinary [][]interface{} `json:"cloudinary"`
}

type UpdateAssetParams struct {
	AssetType         api.AssetType        `json:"-"`
	DeliveryType      api.DeliveryType     `json:"-"`
	PublicID          string               `json:"-"`
	ModerationStatus  api.ModerationStatus `json:"moderation_status,omitempty"`
	RawConvert        string               `json:"raw_convert,omitempty"`
	Ocr               string               `json:"ocr,omitempty"`
	Categorization    string               `json:"categorization,omitempty"`
	Detection         string               `json:"detection,omitempty"`
	SimilaritySearch  string               `json:"similarity_search,omitempty"`
	AutoTagging       float64              `json:"auto_tagging,omitempty"`
	BackgroundRemoval string               `json:"background_removal,omitempty"`
	QualityOverride   int                  `json:"quality_override,omitempty"`
	NotificationUrl   string               `json:"notification_url,omitempty"`
	Tags              api.CldApiArray      `json:"tags,omitempty,omitempty"`
	Context           api.CldApiMap        `json:"context,omitempty"`
	FaceCoordinates   api.Coordinates      `json:"face_coordinates,omitempty"`
	CustomCoordinates api.Coordinates      `json:"custom_coordinates,omitempty"`
	AccessControl     interface{}          `json:"access_control,omitempty"`
}

// UpdateAsset updates details of an existing asset.
//
// Updates one or more of the attributes associated with a specified asset. Note that you can also update
// most attributes of an existing asset using the uploader.Explicit method, which is not rate limited.
//
// https://cloudinary.com/documentation/admin_api#update_details_of_an_existing_resource
func (a *Api) UpdateAsset(ctx context.Context, params UpdateAssetParams) (*AssetResult, error) {
	res := &AssetResult{}
	_, err := a.post(ctx, api.BuildPath(Assets, params.AssetType, params.DeliveryType,
		params.PublicID), params, res)

	return res, err
}