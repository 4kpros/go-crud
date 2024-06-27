# A high-performance and well-structured API for multiple projects

This repository contains a high-performance API written in Go, designed to be reusable and easily integrated into multiple projects. The API is built with a focus on performance, scalability, and maintainability. 

It provides a clean and consistent interface for accessing and managing data, making it an ideal choice for a variety of applications.


# Key features

**- High performance:** The API is designed to handle a high volume of requests with minimal latency.

**- Scalable:** The API can be easily scaled to accommodate increasing workloads.

**- Well-structured:** The API is well-documented and easy to use, with a consistent and intuitive interface.

**- Reusable:** The API can be easily integrated into multiple projects, reducing development time and effort.


# Use cases

**- Web applications:** The API can be used to power web applications that need to access and manage data efficiently.

**- Mobile applications:** The API can be used to develop mobile applications that need to interact with a backend server.


# To get started with the API, follow these steps:

1. Requirements

  - `Redis server`

  - `Memcache server`

  - `Postgres server`

  Others informations such configurations are on `app.env` and `crypto.env` file

2. Clone the repository

 - `git clone https://github.com/your-username/go-api.git`

 - `cd go-api`

3. Install dependencies

  - `go mod download`

4. Run migrations

  - `make build-migrate`

  - `make run-migrate`

5. Run the API

  - `make build`

  - `make run`

  Let's GOOOOOOO 🚀🚀🚀🚀


# Next features

- [ ] 🖋️ Auth:
  - Login (📩Email, 📲Phone number, ☁️Provider['Google', 'Facebook']),
  
  - Register (📩Email, 📲Phone number),
  
  - Activate account,
    
  - Reset password.

- [ ] Users

- [ ] Payments

#Contributing

We welcome contributions to this project. If you have any ideas or improvements, please feel free to open an issue or pull request.

We believe that this API can be a valuable tool for developers who need to build high-performance, scalable, and maintainable applications. We encourage you to try it out and let us know what you think.
