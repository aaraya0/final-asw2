FROM mysql:latest

# Configuración del entorno
ENV MYSQL_ROOT_PASSWORD=aaraya0
ENV MYSQL_DATABASE=users_db
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=aaraya0

# Copiar archivos SQL al contenedor (opcional)
COPY ./database/ /docker-entrypoint-initdb.d/

# Puerto expuesto
EXPOSE 3306

# Comando de inicio
CMD ["mysqld"]
