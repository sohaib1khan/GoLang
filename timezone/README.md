# Time Zone Converter Application

The Time Zone Converter Application is a web-based tool written in Go (Golang). It allows users to:
- View the current time in **EST** (Eastern Standard Time) and **UTC**.
- Enter a time in **EST** to compare it with **UTC**.
- View a table of time zones and their offsets.

---

## Features

- Displays the current time in **EST** and **UTC**.
- Allows time comparisons between EST and UTC.
- Shows a complete table of common time zones with their offsets.
- Responsive and user-friendly web interface.

---

## Directory Structure

Here’s an overview of the file and directory structure of the application:
      
      .
      ├── main.go                   # Main Go application
      ├── static/                   # Static assets (e.g., CSS)
      │   └── styles.css            # CSS for styling the application
      ├── templates/                # HTML templates
      │   └── index.html            # Main HTML file for the application
      ├── timezone-app-linux        # Pre-built Linux executable
      └── timezone-app-windows.exe  # Pre-built Windows executable 




### File Descriptions

- **`main.go`**: The main Go application file containing the server logic.
- **`static/styles.css`**: CSS file that defines the style and layout of the application.
- **`templates/index.html`**: HTML file used as the front-end template for the web application.
- **`timezone-app-linux`**: Pre-built executable for Linux systems.
- **`timezone-app-windows.exe`**: Pre-built executable for Windows systems.

---

## How to Run the Application

### Prerequisites

- **Go** installed on your machine (only if building or running the source code).
- A web browser to access the application.

---

### Running the Pre-Built Executable

1. **For Linux**:
   - Run the following command in your terminal:
     ```bash
     ./timezone-app-linux
     ```
   - Open your browser and navigate to `http://localhost:8080`.

2. **For Windows**:
   - Double-click `timezone-app-windows.exe`.
   - The application will automatically open in your default browser.

---

### Running the Application from Source

1. Clone or download this repository.
2. Navigate to the project directory:
   ```bash
   cd timezone

3. Run the application using:
   ```bash
   go run main.go

4. Open your browser and navigate to http://localhost:8080.

Building the Application

You can build the application for different operating systems using Go’s cross-compilation feature.
Build for Linux

Run the following command:
    ```bash
    GOOS=linux GOARCH=amd64 go build -o timezone-app-linux main.go

Build for Windows

Run the following command:
    ```bash
    GOOS=windows GOARCH=amd64 go build -o timezone-app-windows.exe main.go


How to Use the Application

    View Current Times:
        The home page shows the current time in EST and UTC at the top.

    Compare Times:
        Enter a time in EST (e.g., 02:30 PM) in the comparison form.
        Click "Compare" to see the equivalent time in UTC.

    View Time Zone Table:
        Scroll down to view the table of time zones and their offsets.
