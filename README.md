📘 PRD (Product Requirement Document)
🏥 Product Name
MedBridge
🎯 Objective

Menyediakan sistem backend untuk:
Booking dokter
Manajemen pasien & dokter
Pembayaran appointment
Resep digital

Target:
REST API production-ready
Scalable ke microservices
Clean Architecture
Docker-ready

👥 User Roles
1. Patient
Register
Login
Update profile
Book appointment
Pay appointment
View prescription

2. Doctor
Login
Approve / reject appointment
Create prescription
View schedule

3. Admin (Optional advanced)
View users
Suspend account

📦 Feature Breakdown
🔐 Auth
Register
Login
JWT
Role-based access

👤 User
Update profile
Get detail
List doctors

📅 Appointment
Book
Approve
Cancel
Status tracking
Status flow:
PENDING → APPROVED → COMPLETED
        ↘ REJECTED
💳 Payment
Simulate payment
Webhook endpoint
Payment status:
UNPAID
PAID
FAILED

📄 Prescription
Doctor create prescription
Patient view history

🧠 Non-Functional Requirements
JWT authentication
Graceful shutdown
Context timeout
Transaction handling
Logging
Structured JSON response
Clean architecture
Ready to split microservices