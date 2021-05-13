go build -o bookings cmd/web/*.go
./bookings -dbname=hotel-booking -dbuser=jason.ngan -production=false