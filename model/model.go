package model

type User struct {
	Id        *int    `db:"id"`
	Name      string  `db:"name"`
	Email     string  `db:"email"`
	Password  string  `db:"password"`
	DeletedAt *string `db:"deleted_at"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}

type Size struct {
	Id        *int    `db:"id"`
	Name      string  `db:"name"`
	DeletedAt *string `db:"deleted_at"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}

type Color struct {
	Id             *int    `db:"id"`
	Name           string  `db:"name"`
	SecondaryColor *string `db:"secondary_color"`
	DeletedAt      *string `db:"deleted_at"`
	CreatedAt      string  `db:"created_at"`
	UpdatedAt      string  `db:"updated_at"`
}

type Category struct {
	Id        *int    `db:"id"`
	Name      string  `db:"name"`
	Slug      string  `db:"slug"`
	Index     *int    `db:"index"`
	DeletedAt *string `db:"deleted_at"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}

type SubCategory struct {
	Id          *int    `db:"id"`
	CategoryId  *int    `db:"category_id"`
	Name        string  `db:"name"`
	Slug        string  `db:"slug"`
	Description *string `db:"description"`
	DeletedAt   *string `db:"deleted_at"`
	UpdatedAt   string  `db:"updated_at"`
	CreatedAt   string  `db:"created_at"`
}

type Product struct {
	Id        *int    `db:"id"`
	Name      string  `db:"name"`
	Weight    float32 `db:"weight"`
	Width     float32 `db:"width"`
	Height    float32 `db:"height"`
	Price     float32 `db:"price"`
	DeletedAt *string `db:"deleted_at"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}

type ProductVariation struct {
	Id        *int    `db:"id"`
	ProductId int     `db:"product_id"`
	SizeId    int     `db:"size_id"`
	ColorId   *int    `db:"color_id"`
	DeletedAt *string `db:"deleted_at"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
}

type ProductCategory struct {
	Id            *int    `db:"id"`
	ProductId     int     `db:"product_id"`
	SubCategoryId int     `db:"sub_category_id"`
	DeletedAt     *string `db:"deleted_at"`
	CreatedAt     string  `db:"created_at"`
	UpdatedAt     string  `db:"updated_at"`
}

type ProductStock struct {
	Id                 *int    `db:"id"`
	ProductVariationId int     `db:"product_variation_id"`
	Quantity           int     `db:"quantity"`
	DeletedAt          *string `db:"deleted_at"`
	CreatedAt          string  `db:"created_at"`
	UpdatedAt          string  `db:"updated_at"`
}
