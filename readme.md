# Setup DB
```sql
CREATE DATABASE dotdb
    DEFAULT CHARACTER SET = 'utf8mb4';
```

# Environment variable
```bash
cp .running.sh.example .running.sh
```
* on file running.sh, don't forget to change the user and password with your local settings


# Running
```bash
sh running.sh
```

## OR withh no `running.sh` file

```bash 
export MYSQL_URL="<your-user>:<your-password>@tcp(127.0.0.1:3306)/dotdb?charset=utf8mb4&parseTime=True&loc=Local"
go run .
```

# Tech Stack
* [goriber.io](https://gofiber.io/)
* [MySQL](https://www.mysql.com/)
* [gorm](https://gorm.io/)

# Fitur
* Automatically migrate table and column

# Endpoints

## Employee