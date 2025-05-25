# CoworkingRentalSystem

A microservice-based backend system for managing coworking space rentals. It allows users to register, view available spaces, create and cancel bookings, and receive confirmation emails. The project is structured with clean architecture and includes gRPC communication, message queues, caching, and email notifications.

---

## 📌 Project Overview

This project is a backend for a coworking space booking platform. It includes the following microservices:

- **User Service** – handles user registration, login, and profile management.
- **Space Service** – manages coworking space CRUD operations.
- **Booking Service** – handles bookings (create, list, cancel, get by ID).
- **API Gateway** – single entry point to communicate with all services.

---

## ⚙️ Technologies Used

- **Go (Golang)** – core language
- **gRPC** – communication between services
- **MongoDB** – database for all services
- **Redis** – caching for bookings
- **NATS** – message queue for async events
- **SMTP (Gmail)** – sending email confirmations
- **Docker & Docker Compose** – containerization and orchestration

---

## 🚀 How to Run Locally

### With Docker:

```bash
docker compose build
docker compose up
```
---
## 🧪 How to Run Tests
Tests planned to be added soon.

---

## 📡 gRPC Endpoints

## 👤 UserService (port 50051)

- RegisterUser(email, password, full_name)
- LoginUser(email, password)
- GetUserProfile(user_id)

## 🏢 SpaceService (port 50052)
- CreateSpace(...)

- UpdateSpace(...)
- DeleteSpace(space_id)
- GetSpaceByID(space_id)
- ListSpaces()

## 📦BookingService (port 50053)
- CreateBooking(booking_id, user_id, space_id, date, email)
- ListBookings(user_id)
- GetBookingByID(booking_id)
- CancelBooking(booking_id)

---
## ✅ Features Implemented
- Clean architecture (entities, usecase, repository, transport)
- gRPC APIs for user, space, and booking 
- MongoDB integration in all services 
- Redis caching in booking service 
- NATS message queue (booking.created event)
- Email notification system (Gmail SMTP)
- Dockerized services via Docker Compose 
- API Gateway with HTTP → gRPC routing 
- CancelBooking functionality 
- GetBookingByID functionality 
- Token-based authentication (planned)
- Unit & integration tests (planned)

---
## 📬 Authors
- Bolatkan Yerassyl (yerassylv)
- Salimgerey Yerassyl (Kampo77)
 

