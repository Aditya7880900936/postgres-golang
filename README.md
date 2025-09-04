# Postgres + Go REST API

A simple RESTful API built with **Go**, **Gorilla Mux**, and **PostgreSQL**. This project provides a foundation to **create, read, update, and delete** records from a stock database.

The project is fully containerized using **Docker Compose**, which sets up both the Go application and the PostgreSQL database.

---

## Folder Structure

```
.
├── controllers
├── db
│   └── init.sql
├── middlewares
│   └── middlewares.go
├── migrations
├── models
│   └── models.go
├── routers
│   └── router.go
├── .env
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
└── main.go
```

---

## Prerequisites

- **Docker & Docker Compose**
- **Go 1.20+** (only required if you choose to run without Docker)
- An API testing tool like **Postman**

---

## 🚀 Setup & Run Locally

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Aditya7880900936/postgres-golang.git
    cd postgres-golang
    ```

2.  **Create a `.env` file** in the root directory with your database configuration:
    ```env
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=yourpassword
    POSTGRES_DB=stocksdb
    POSTGRES_PORT=5432
    POSTGRES_URL=postgres://postgres:yourpassword@postgres:5432/stocksdb?sslmode=disable
    ```

3.  **Run with Docker Compose:**
    This command will build the Go application image and start the `go-app` and `postgres` services.
    ```bash
    docker-compose up --build
    ```

Your API will be live and accessible at `http://localhost:8080`.

---

## 📦 Stock Model

All API requests and responses for stock data use the following JSON structure.

**JSON Structure:**
```json
{
  "stockid": 1,
  "stockname": "Apple",
  "stockprice": 1500.0,
  "stockcompany": "Apple Inc"
}
```

**Fields:**
- `stockid` (int64): The unique identifier for the stock.
- `stockname` (string): The name of the stock.
- `stockprice` (float64): The current price of the stock.
- `stockcompany` (string): The parent company of the stock.

---

## Routes

| Method   | Route                   | Description              |
| -------- | ----------------------- | ------------------------ |
| `GET`    | `/api/stock`            | Get all stocks           |
| `GET`    | `/api/stock/{id}`       | Get a single stock by ID |
| `POST`   | `/api/newstock`         | Create a new stock       |
| `PUT`    | `/api/stock/{id}`       | Update a stock by ID     |
| `DELETE` | `/api/deletestock/{id}` | Delete a stock by ID     |

---

## 🧪 Testing with Postman

#### Get All Stocks
- **Method**: `GET`
- **URL**: `http://localhost:8080/api/stock`

#### Get Stock by ID
- **Method**: `GET`
- **URL**: `http://localhost:8080/api/stock/1`

#### Create a New Stock
- **Method**: `POST`
- **URL**: `http://localhost:8080/api/newstock`
- **Body** (JSON):
  ```json
  {
    "stockname": "Apple",
    "stockprice": 1500,
    "stockcompany": "Apple Inc"
  }
  ```

#### Update Stock
- **Method**: `PUT`
- **URL**: `http://localhost:8080/api/stock/1`
- **Body** (JSON):
  ```json
  {
    "stockname": "Apple Inc",
    "stockprice": 1550,
    "stockcompany": "Apple Inc"
  }
  ```

#### Delete Stock
- **Method**: `DELETE`
- **URL**: `http://localhost:8080/api/deletestock/1`

---

## Migrations

- All database migration files are located in the `/migrations` folder.
- Migrations are applied automatically when the application starts up with `docker-compose up`.

---

## Notes

- Ensure the Postgres container is fully up and healthy before the Go application attempts to connect.
- Environment variables are loaded from the `.env` file at runtime.
- All API responses are in JSON format.

---

## Author

- **Aditya Sanskar Srivastav**
