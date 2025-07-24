/*
  - Product csv columns
    timestamp TEXT,
    title TEXT,
    seller_name TEXT,
    brand TEXT,
    description TEXT,
    initial_price TEXT,
    final_price TEXT,
    currency TEXT,
    availability TEXT,
    reviews_count TEXT,
    categories TEXT,
    asin TEXT,
    buybox_seller TEXT,
    number_of_sellers TEXT,
    root_bs_rank TEXT,
    answered_questions TEXT,
    domain TEXT,
    images_count TEXT,
    url TEXT,
    video_count TEXT,
    image_url TEXT,
    item_weight TEXT,
    rating TEXT,
    product_dimensions TEXT,
    seller_id TEXT,
    date_first_available TEXT,
    discount TEXT,
    model_number TEXT,
    manufacturer TEXT,
    department TEXT,
    plus_content TEXT,
    upc TEXT,
    video TEXT,
    top_review TEXT,
    variations TEXT,
    delivery TEXT,
    features TEXT,
    format TEXT,
    buybox_prices TEXT,
    parent_asin TEXT,
    input_asin TEXT,
    ingredients TEXT,
    origin_url TEXT,
    bought_past_month TEXT,
    is_available TEXT,
    root_bs_category TEXT,
    bs_category TEXT,
    bs_rank TEXT,
    badge TEXT,
    subcategory_rank TEXT,
    amazon_choice TEXT,
    images TEXT,
    product_details TEXT,
    prices_breakdown TEXT,
    country_of_origin TEXT *
*/
package model


type Product struct {
	Timestamp          string `json:"timestamp"`
	Title              string `json:"title"`
	SellerName         string `json:"seller_name"`
	Brand              string `json:"brand"`
	Description       string `json:"description"`
	InitialPrice       string `json:"initial_price"`
	FinalPrice         string `json:"final_price"`
	Currency           string `json:"currency"`
	Availability       string `json:"availability"`
	ReviewsCount       string `json:"reviews_count"`
	Categories         string `json:"categories"`
	Asin               string `json:"asin"`
	BuyboxSeller       string `json:"buybox_seller"`
	NumberOfSellers    string `json:"number_of_sellers"`
	RootBsRank         string `json:"root_bs_rank"`
	AnsweredQuestions  string `json:"answered_questions"`	
	Domain             string `json:"domain"`
	ImagesCount        string `json:"images_count"`
	URL                string `json:"url"`
	VideoCount         string `json:"video_count"`
	ImageURL           string `json:"image_url"`
	ItemWeight         string `json:"item_weight"`
	Rating             string `json:"rating"`
	ProductDimensions  string `json:"product_dimensions"`
	SellerID           string `json:"seller_id"`
	DateFirstAvailable string `json:"date_first_available"`
	Discount           string `json:"discount"`
	ModelNumber        string `json:"model_number"`
	Manufacturer       string `json:"manufacturer"`
	Department         string `json:"department"`
	PlusContent        string `json:"plus_content"`
	UPC                string `json:"upc"`
	Video              string `json:"video"`
	TopReview          string `json:"top_review"`
	Variations         string `json:"variations"`
	Delivery           string `json:"delivery"`
	Features           string `json:"features"`
	Format             string `json:"format"`
	BuyboxPrices       string `json:"buybox_prices"`
	ParentAsin         string `json:"parent_asin"`
	InputAsin          string `json:"input_asin"`
	Ingredients        string `json:"ingredients"`
	OriginURL          string `json:"origin_url"`
	BoughtPastMonth    string `json:"bought_past_month"`
	IsAvailable        string `json:"is_available"`
	RootBsCategory     string `json:"root_bs_category"`
	BsCategory         string `json:"bs_category"`
	BsRank             string `json:"bs_rank"`
	Badge              string `json:"badge"`
	SubcategoryRank    string `json:"subcategory_rank"`
	AmazonChoice       string `json:"amazon_choice"`
	Images             string `json:"images"`
	ProductDetails     string `json:"product_details"`
	PricesBreakdown    string `json:"prices_breakdown"`
	CountryOfOrigin    string `json:"country_of_origin"`
}

/**
{
  "id": "12345",
  "name": "Wireless Bluetooth Headphones",
  "description": "Over-ear noise-cancelling Bluetooth headphones with 30 hours of battery life.",
  "price": 89.99,
  "currency": "USD",
  "stock": 150,
  "category": "electronics",
  "brand": "SoundPro",
  "images": [
    {
      "url": "https://example.com/images/headphones1.jpg",
      "alt": "Wireless Bluetooth Headphones - Front View"
    },
    {
      "url": "https://example.com/images/headphones2.jpg",
      "alt": "Wireless Bluetooth Headphones - Side View"
    }
  ],
  "variants": [
    {
      "color": "Black",
      "sku": "12345-BLK",
      "stock": 100
    },
    {
      "color": "White",
      "sku": "12345-WHT",
      "stock": 50
    }
  ],
  "ratings": {
    "average": 4.5,
    "count": 324
  }
}
**/

type TrendyolProduct struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	Currency    string `json:"currency"`
	Stock       string    `json:"stock"`
	Category    string `json:"category"`
	Brand       string `json:"brand"`
	Images      []string `json:"images"`
	Variants    []string `json:"variants"`//list model
	Ratings     string `json:"ratings"`
}