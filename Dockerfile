# Use an official PostgreSQL image as a parent image
FROM postgres:13

# Set environment variables (replace with your own values)
ENV POSTGRES_DB=mybankdb
ENV POSTGRES_USER=myuser
ENV POSTGRES_PASSWORD=mypassword
