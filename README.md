<p align="center">
<h1 align="center">MyFitnessPal-Grafana</h1>
  <p align="center">
    A simple App which visualizes your MyFitnessPal exported csv data using Grafana.
    <br/>
  </p>
</p>


## Project Demo
https://github.com/Tejasmadhukar/MyFitnessPal-Grafana/assets/39000441/c10d0c09-4079-4327-8c7a-b37f093577b9

## Getting Started

Make sure you have docker installed.

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/Tejasmadhukar/MyFitnessPal-Grafana
   ```
2. Run grafana-instance and go-server
   ```sh
   docker-compose up
   ```
3. Go to http://localhost:3000 to your grafana app and [follow these steps](https://grafana.com/docs/grafana/latest/administration/service-accounts/#create-a-service-account-in-grafana) to create a service account.
4. copy .env.example to .env and paste the service token.
5. restart both services with docker
   ```sh
   docker-compose up
   ```
6. Use the app at http://localhost:80 !!!

### Tech Stack 
1. go & [go-chi](https://go-chi.io/#/)
2. [htmx](https://htmx.org/)
3. [grafana](https://grafana.com/)
