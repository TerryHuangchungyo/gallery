FROM mysql:8.0.21
COPY ./gallery_backup.sql /backup/gallery_backup.sql
COPY ./init_gallery.sh /docker-entrypoint-initdb.d/init_gallery.sh
ENV MYSQL_ROOT_PASSWORD root
