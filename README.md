# D-Dashboard

D-Dashboard is a versatile system metric monitoring application built using React and Go. This application is designed to be independent of any operating system, providing real-time insights into system performance metrics.

Features


- Real-time Monitoring: Track system metrics in real time.
- Cross-Platform Compatibility: Works seamlessly on any operating system.
- User-Friendly Interface: Simple and intuitive UI built with React.
- Robust Backend: Efficient data handling with a Go backend.

Table of Contents


- #features
- #installation
- #usage
- #configuration

Installation


Prerequisites

- Docker
- Node.js
- Go

Steps

1. Clone the Repository:

bash
```
git clone (link unavailable)
cd D-Dashboard
```

2. Build and Run with Docker:

bash
```
docker-compose up --build
```

This command will build and start the application using Docker, ensuring it runs independently of the host operating system.

### Access the Application

Open your browser and navigate to http://localhost:3000 to access the D-Dashboard interface.

Monitor Metrics

The dashboard provides an overview of system metrics:

- CPU Usage: View real-time CPU utilization.
- Memory Usage: Check current memory consumption.
- Disk Usage: Monitor disk space usage.
- Network Activity: Track network input/output.
- Refresh Rate: The dashboard updates metrics every 5 seconds, providing near real-time data.

Configuration


Frontend (React)

Located in the /frontend directory.

- Install Dependencies:

bash
```
cd frontend
npm install
```

- Start the Development Server:

bash
```
npm start
```

### Backend (Go)

Located in the `/backend` directory.

* Install Dependencies:
bash
```
cd backend
go get ./...
```
- Run the Backend Server:

bash
```
go run main.go
```
