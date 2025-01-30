# Veterinary Clinic API

This API allows managing pet records in a veterinary clinic.

## 🛠 Technologies Used
- **Golang** (with `gorilla/mux`)
- **GORM** (for database integration)
- **Google Cloud Run** (for deploy in GCP)
- **Cloud SQL** (SQL instance in GCP)
- **Secret Manager** (Managing secrets, like the database and Google Cloud informations)
- **Docker** 



This structure follows an architecture pattern known as **MVC (Model-View-Controller)**

- **config**: Stores database configurations.
- **controllers**: Contains controllers that handle HTTP requests and interact with services and models.
- **model**: Contains data model
- **routes**: Manages routing of requests to the appropriate controllers.
- **services**: Contains business logic, which can be seen as a **Service** or **Business Logic** layer.


```json
📂 config
 ├── 📄 database.go
📂 controllers
 ├── 📄 pet.go
📂 model
 ├── 📄 pet.go
📂 routes
 ├── 📄 routes.go
📂 services
 ├── 📄 pet_service.go
📄 Dockerfile
📄 go.mod
📄 go.sum
📄 main.go
```



## 📌 Endpoints

### 🔹 Get all pets
`GET /pets`
#### **Success Response (200 OK)**
```json
[
  {
    "ID": 1,
    "Name": "Buddy",
    "Age": 3,
    "Species": "Dog",
    "Owner": "John Doe",
    "CreatedAt": "2024-01-01T12:00:00Z",
    "UpdatedAt": "2024-01-02T12:00:00Z"
  }
]
```

---

### 🔹 Get pet by ID
`GET /pets/{id}`
#### **Success Response (200 OK)**
```json
{
  "ID": 1,
  "Name": "Buddy",
  "Age": 3,
  "Species": "Dog",
  "Owner": "John Doe",
  "CreatedAt": "2024-01-01T12:00:00Z",
  "UpdatedAt": "2024-01-02T12:00:00Z"
}
```
#### **Error (404 Not Found)**
```json
{
  "error": "Pet not found"
}
```

---

### 🔹 Create a new pet
`POST /pets`
#### **Request Body**
```json
{
  "Name": "Charlie",
  "Age": 2,
  "Species": "Cat",
  "Owner": "Jane Doe"
}
```
#### **Success Response (201 Created)**
```json
{
  "ID": 2,
  "Name": "Charlie",
  "Age": 2,
  "Species": "Cat",
  "Owner": "Jane Doe",
  "CreatedAt": "2024-01-03T12:00:00Z",
  "UpdatedAt": "2024-01-03T12:00:00Z"
}
```
#### **Error (400 Bad Request)**
```json
{
  "error": "Invalid data format"
}
```

---

### 🔹 Update a pet
`PUT /pets/{id}`
#### **Request Body** (Can contain one or more fields for update)
```json
{
  "Age": 4,
  "Owner": "Michael Scott"
}
```
#### **Success Response (200 OK)**
```json
{
  "ID": 1,
  "Name": "Buddy",
  "Age": 4,
  "Species": "Dog",
  "Owner": "Michael Scott",
  "CreatedAt": "2024-01-01T12:00:00Z",
  "UpdatedAt": "2024-01-04T12:00:00Z"
}
```
#### **Error (404 Not Found)**
```json
{
  "error": "Pet not found"
}
```

---

### 🔹 Delete a pet
`DELETE /pets/{id}`
#### **Success Response (200 OK)**
```json
{
  "ID": 1,
  "Name": "Buddy",
  "Age": 4,
  "Species": "Dog",
  "Owner": "Michael Scott"
}
```
#### **Error (404 Not Found)**
```json
{
  "error": "Pet not found"
}
```

---

## 🏗 Pet JSON Structure
```json
{
  "ID": 1,
  "Name": "Buddy",
  "Age": 3,
  "Species": "Dog",
  "Owner": "John Doe",
  "CreatedAt": "2024-01-01T12:00:00Z",
  "UpdatedAt": "2024-01-02T12:00:00Z"
}
```
