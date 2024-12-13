# Go Dashboard Project

## **Overview**

The Go Dashboard is a lightweight web-based application built using Go (Golang) and Fiber, designed to manage bookmarks and categories. It provides an intuitive UI and a persistent data store, making it easy to organize and retrieve your bookmarks.

The app is portable and can run on Windows, Linux, or any platform supporting Go. Additionally, it can be packaged as a Docker container for convenient deployment. The focus is on simplicity, portability, and providing a highly responsive and visually appealing interface for users. With the ability to maintain data persistently and handle various environments, this app is ideal for personal use or as a foundation for more complex applications.

---

## **Features**

1. **Bookmark Management**:
   - Add, edit, and delete bookmarks.
   - Organize bookmarks by categories, ensuring easy navigation and organization.
2. **Category Management**:
   - Add and remove categories dynamically.
   - Ensure categories with associated bookmarks cannot be deleted to prevent data inconsistency.
3. **Persistent Data**:
   - Uses a `data.json` file to store bookmarks and categories persistently.
   - Ensures data remains intact even after application restarts.
4. **Responsive UI**:
   - Built with HTML, CSS, and JavaScript for a sleek and interactive user experience.
   - Features animations like floating effects and a northern lights background, enhancing the visual appeal.
5. **Dockerized**:
   - Comes with a `Dockerfile` and `docker-compose.yml` for easy containerization and deployment.
6. **Portability**:
   - Runs seamlessly on multiple platforms including Windows, Linux, and macOS.

---

## **Technologies Used**

### **Backend**

- **Go (Golang)**:
  - Core programming language used for the application.
  - Fiber framework provides a fast and efficient HTTP server.
- **Fiber**:
  - A fast, lightweight web framework inspired by Express.js.
  - Used for routing, serving static files, and rendering templates effectively.

### **Frontend**

- **HTML**: Defines the structure of the web interface.
- **CSS**: Responsible for styling, animations, and creating an engaging UI.
- **JavaScript**: Handles interactivity, dynamic updates, and seamless communication with the backend.

### **Data Storage**

- **JSON File**:
  - Provides persistent storage for bookmarks and categories.
  - Easy to manage and modify for small-scale applications.

### **Containerization**

- **Docker**:
  - Packages the app into a container for easy distribution.
  - Ensures consistency across different environments.
- **Docker Compose**:
  - Simplifies multi-container application management.

---

## **Setup Instructions**

### **Requirements**

- Go 1.21 or higher.
- Docker (optional for containerization).
- Node.js and Electron (optional for creating a desktop app).

### **Steps to Run Locally**

#### 1. **Clone the Repository**

```bash
# Replace <repo-url> with the actual repository URL
git clone <repo-url>
cd Go-Dashboard
```

#### 2. **Build and Run the App**

```bash
# Build the Go application
go build -o GoDashboard main.go

# Run the application
./GoDashboard
```

#### 3. **Access the App**

Open your browser and navigate to:

```
http://localhost:3000
```

### **Run with Docker**

#### 1. **Build the Docker Image**

```bash
docker-compose build
```

#### 2. **Run the Application**

```bash
docker-compose up
```

#### 3. **Access the App**

Open your browser and navigate to:

```
http://localhost:3000
```

---

## **File Structure**

```plaintext
Go-Dashboard/
├── data.json                # Persistent storage for bookmarks and categories
├── Dockerfile               # Docker instructions for building the image
├── docker-compose.yml       # Docker Compose for multi-container deployment
├── go.mod                   # Go module dependencies
├── go.sum                   # Go module checksums
├── main.go                  # Main application logic
├── static/                  # Static assets (CSS, JavaScript)
│   ├── css/
│   │   └── styles.css       # Stylesheet for UI
│   └── js/
│       └── scripts.js       # Client-side JavaScript
├── vendor/                  # Go dependencies (for offline builds)
└── views/
    └── index.html           # HTML template for the app
```

---

## **Detailed Explanations**

### **Go (main.go)**

- The `main.go` file initializes the server using Fiber.
- Routes:
  - `/` - Serves the homepage.
  - `/api/data` (GET) - Fetches bookmarks and categories.
  - `/api/data` (POST) - Updates bookmarks and categories.
- Persistent storage is handled via `data.json`, ensuring data durability and integrity.

### **Frontend (HTML, CSS, JavaScript)**

- **index.html**:
  - Provides the structure for the web interface.
- **styles.css**:
  - Applies a northern lights background, floating animations, and other visual enhancements.
- **scripts.js**:
  - Handles dynamic updates, bookmark and category management, and communication with the backend using Fetch API.

### **Docker**

- **Dockerfile**:
  - Uses a multi-stage build process to create small, optimized containers.
  - Includes static files and the Go binary for efficient deployment.
- **docker-compose.yml**:
  - Simplifies the setup process by managing the container's lifecycle and dependencies.

---

---

##

