package logic

import (
	"database/sql"
	"fmt"
	"time"
	"warehouse/models"
	"warehouse/utils"
)

// ProductLogic holds configuration for product logic.
type ProductLogic struct{ Config utils.Config }

// validateProduct validates the fields of a product before it's added or updated.
func validateProduct(product *models.Product) error {
	if product.Name == "" {
		return fmt.Errorf("product name is required")
	}

	if product.SKU == "" {
		return fmt.Errorf("SKU is required")
	}

	if product.Qty <= 0 {
		return fmt.Errorf("quantity must be greater than zero")
	}

	if product.Location == "" {
		return fmt.Errorf("location is required")
	}
	return nil
}

// AddProduct adds a new product to database.
func (pl *ProductLogic) AddProduct(product *models.Product) (string, error) {
	// Ensure a valid database connection.
	if err := utils.EnsureDBConnection(pl.Config); err != nil {
		return "", fmt.Errorf("failed ensure database connection: %w", err)
	}

	// Validate product details before insertion.
	if err := validateProduct(product); err != nil {
		return "", err
	}

	// Insert product into database.
	_, err := utils.DB.Exec("INSERT INTO products (name, sku, qty, location, status) VALUES (?,?,?,?,?)",
		product.Name, product.SKU, product.Qty, product.Location, product.Status)

	if err != nil {
		return "", fmt.Errorf("failed insert product into database: %w", err)
	}
	return "Product added successfully", nil
}

// GetProducts retrieves all products from database.
func (pl *ProductLogic) GetProducts() ([]models.Product, error) {
	// Ensure a valid database connection.
	if err := utils.EnsureDBConnection(pl.Config); err != nil {
		return nil, fmt.Errorf("failed ensure database connection: %w", err)
	}

	// Query get all products.
	rws, err := utils.DB.Query("SELECT id, name, sku, qty, location, status, created_at, updated_at FROM products")
	if err != nil {
		return nil, fmt.Errorf("failed query products: %w", err)
	}
	defer rws.Close()

	// Process results into a slice of Product models.
	products := []models.Product{}
	for rws.Next() {
		var createdAt, updatedAt string
		product := models.Product{}

		// Scan database row into product struct.
		if err := rws.Scan(&product.ID, &product.Name, &product.SKU, &product.Qty,
			&product.Location, &product.Status, &createdAt, &updatedAt); err != nil {
			return nil, fmt.Errorf("failed scan product: %w", err)
		}

		// Parse created_at and updated_at timestamps into time.Time.
		if t, err := time.Parse("2006-01-02 15:04:05", createdAt); err == nil {
			product.CreatedAt = t
		} else {
			return nil, fmt.Errorf("failed parse created_at: %w", err)
		}
		if t, err := time.Parse("2006-01-02 15:04:05", updatedAt); err == nil {
			product.UpdatedAt = t
		} else {
			return nil, fmt.Errorf("failed parse updated_at: %w", err)
		}

		// Append product to the result slice.
		products = append(products, product)
	}
	if err := rws.Err(); err != nil {
		return nil, fmt.Errorf("error rows: %w", err)
	}

	return products, nil
}

// GetProductByID retrieves a single product by its ID.
func (pl *ProductLogic) GetProductByID(id int) (*models.Product, error) {
	// Ensure a valid database connection.
	if err := utils.EnsureDBConnection(pl.Config); err != nil {
		return nil, fmt.Errorf("failed ensure database connection: %w", err)
	}

	product := models.Product{}
	var createdAts, updatedAts string
	// Query to get the product by ID.
	err := utils.DB.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(
		&product.ID, &product.Name, &product.SKU, &product.Qty, &product.Location, &product.Status, &createdAts, &updatedAts)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("failed get products: %w", err)
		}
		return nil, fmt.Errorf("failed query products: %w", err)
	}

	// Parse created_at and updated_at timestamps into time.Time.
	if t, err := time.Parse("2006-01-02 15:04:05", createdAts); err == nil {
		product.CreatedAt = t
	} else {
		return nil, fmt.Errorf("failed parse created_at: %w", err)
	}
	if t, err := time.Parse("2006-01-02 15:04:05", updatedAts); err == nil {
		product.UpdatedAt = t
	} else {
		return nil, fmt.Errorf("failed parse updated_at: %w", err)
	}

	return &product, nil
}

// UpdateProduct updates an existing product in database.
func (pl *ProductLogic) UpdateProduct(product *models.Product) (string, error) {
	// Ensure a valid database connection.
	if err := utils.EnsureDBConnection(pl.Config); err != nil {
		return "", fmt.Errorf("failed ensure database connection: %w", err)
	}

	// Validate product details before update.
	if err := validateProduct(product); err != nil {
		return "", err
	}

	product.UpdatedAt = time.Now()

	// Update product in database.
	_, err := utils.DB.Exec("UPDATE products SET name = ?, sku = ?, qty = ?, location = ?, status = ?, updated_at = ? WHERE id = ?",
		product.Name, product.SKU, product.Qty, product.Location, product.Status, product.UpdatedAt, product.ID)
	if err != nil {

		return "", fmt.Errorf("failed update product: " + err.Error())
	}

	return "Product updated successfully", nil
}

// DeleteProduct deletes a product from the database by its ID.
func (pl *ProductLogic) DeleteProduct(id int) (string, error) {
	// Delete product from database.
	_, err := utils.DB.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return "", fmt.Errorf("failed to delete product: " + err.Error())
	}

	return "Product deleted successfully", nil
}
