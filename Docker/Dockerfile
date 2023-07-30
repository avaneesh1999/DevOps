##--------------------Sample Docker file----------------

# Use an official web server as the base image (for example, Nginx)
FROM nginx:latest

# Copy the static website files to the web server's root directory
COPY ./website /usr/share/nginx/html

# Expose the port on which the web server will listen (usually port 80 for HTTP)
EXPOSE 80

# Define the command to start the web server (handled by the base Nginx image)
CMD ["nginx", "-g", "daemon off;"]
