# Hospital Management System API

## Description
This is a simple API for a hospital management system deployed on Railway. It is built using Go and Gin. The database used is PlanetScale/MySQL. The API is deployed on Railway and can be accessed [here](https://hospital-api.up.railway.app/patients).

## Documentation
The documentation for this API can be found [here](https://documenter.getpostman.com/view/26210548/2s9YkgEkcX).

## Database Schema
The database schema for this API can be found [here](https://dbdiagram.io/d/64dd534e02bd1c4a5ee6ab86).

## Request Methods
- `GET` - Used to retrieve data from the server.
- `POST` - Used to send data to the server to create a resource.
- `PUT` - Used to send data to the server to update a resource.
- `DELETE` - Used to delete a resource from the server.

## Status Codes
- `200` - OK. The request has succeeded.
- `400` - Bad Request. The server could not understand the request due to invalid syntax.
- `404` - Not Found. The server can not find the requested resource.

## Endpoints

### Patients
- `GET /patients` - Get all patients
- `GET /patients/:id` - Get a patient by id
- `POST /patients` - Create a new patient
- `PUT /patients/:id` - Update a patient by id
- `DELETE /patients/:id` - Delete a patient by id

### Doctors
- `GET /doctors` - Get all doctors
- `GET /doctors/:id` - Get a doctor by id
- `POST /doctors` - Create a new doctor
- `PUT /doctors/:id` - Update a doctor by id
- `DELETE /doctors/:id` - Delete a doctor by id

### Nurses
- `GET /nurses` - Get all nurses
- `GET /nurses/:id` - Get a nurse by id
- `POST /nurses` - Create a new nurse
- `PUT /nurses/:id` - Update a nurse by id
- `DELETE /nurses/:id` - Delete a nurse by id

### Wardboys
- `GET /wardboys` - Get all wardboys
- `GET /wardboys/:id` - Get a wardboy by id
- `POST /wardboys` - Create a new wardboy
- `PUT /wardboys/:id` - Update a wardboy by id
- `DELETE /wardboys/:id` - Delete a wardboy by id

### Rooms
- `GET /rooms` - Get all rooms
- `GET /rooms/:id` - Get a room by id
- `POST /rooms` - Create a new room
- `PUT /rooms/:id` - Update a room by id
- `DELETE /rooms/:id` - Delete a room by id

### Treatments
- `GET /treatments` - Get all treatments
- `GET /treatments/:id` - Get a treatment by id
- `POST /treatments` - Create a new treatment
- `PUT /treatments/:id` - Update a treatment by id
- `DELETE /treatments/:id` - Delete a treatment by id

### Numbers
- `GET /numbers` - Get all numbers
- `GET /numbers/:id` - Get a number by id
- `POST /numbers` - Create a new number
- `PUT /numbers/:id` - Update a number by id
- `DELETE /numbers/:id` - Delete a number by id

### Bills
- `GET /bills` - Get all bills
- `GET /bills/:id` - Get a bill by id
- `POST /bills` - Create a new bill
- `PUT /bills/:id` - Update a bill by id
- `DELETE /bills/:id` - Delete a bill by id





