# !/bin/.sh

until mysqladmin ping -h db --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo " go app is started!"
exec go run main.go