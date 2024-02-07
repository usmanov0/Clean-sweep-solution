CREATE TABLE "cart"(
    "id" SERIAL PRIMARY KEY,
    "user_id" INTEGER REFERENCES "users"("id"),
    "status" BOOLEAN
);

CREATE TABLE "cart_items"(
    "id" SERIAL PRIMARY KEY,
    "cart_id" INTEGER REFERENCES "cart"("id"),
    "product_id" INTEGER REFERENCES "products"("id"),
    "quantity" INTEGER
);