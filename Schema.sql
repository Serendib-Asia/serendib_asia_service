This is my database, please check dtos created correctly or not if is no match with sql fix it 

-- ==============================
-- 🔹 MASTER LOOKUP TABLES
-- ==============================

CREATE TABLE purpose_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) UNIQUE NOT NULL -- e.g., Sell, Rent, Stay
);

CREATE TABLE property_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL -- e.g., Apartment, House, Villa, Land
);

CREATE TABLE furniture_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL -- e.g., Fully Furnished, Semi Furnished
);

CREATE TABLE property_conditions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL -- e.g., New, Used, Renovated
);

CREATE TABLE utilities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL -- e.g., Electricity, Water, etc.
);

CREATE TABLE amenities (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL -- e.g., Schools, Gym, Playground
);

-- ==============================
-- 🔹 USERS
-- ==============================

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    phone_number VARCHAR(15),
    profile_image TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ==============================
-- 🔹 PROPERTIES (Final Version)
-- ==============================

CREATE TABLE properties (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    title VARCHAR(150),
    description TEXT,
    purpose_id INTEGER REFERENCES purpose_types(id),
    property_type_id INTEGER REFERENCES property_types(id),
    furniture_type_id INTEGER REFERENCES furniture_types(id),
    condition_id INTEGER REFERENCES property_conditions(id),
    bedrooms INTEGER,
    bathrooms INTEGER,
    size FLOAT,
    size_unit VARCHAR(20), -- e.g., Sqft, Perch
    city VARCHAR(50),
    address TEXT,
    postal_code VARCHAR(10),
    latitude FLOAT,
    longitude FLOAT,
    price FLOAT,
    price_unit VARCHAR(20), -- e.g., LKR
    is_negotiable BOOLEAN DEFAULT FALSE,
    rental_period VARCHAR(20), -- Monthly, Weekly, etc.
    is_refundable BOOLEAN DEFAULT FALSE,
    pricing_type VARCHAR(10) CHECK (pricing_type IN ('sell', 'rent', 'stay')) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- ==============================
-- 🔹 MANY-TO-MANY RELATIONS
-- ==============================

CREATE TABLE property_amenities (
    property_id INTEGER REFERENCES properties(id),
    amenity_id INTEGER REFERENCES amenities(id),
    PRIMARY KEY (property_id, amenity_id)
);

CREATE TABLE property_utilities (
    property_id INTEGER REFERENCES properties(id),
    utility_id INTEGER REFERENCES utilities(id),
    PRIMARY KEY (property_id, utility_id)
);

-- ==============================
-- 🔹 MEDIA & FAVORITES
-- ==============================

CREATE TABLE property_images (
    id SERIAL PRIMARY KEY,
    property_id INTEGER REFERENCES properties(id),
    url TEXT NOT NULL,
    is_primary BOOLEAN DEFAULT FALSE
);

CREATE TABLE favourites (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id),
    property_id INTEGER REFERENCES properties(id),
    UNIQUE(user_id, property_id)
);
