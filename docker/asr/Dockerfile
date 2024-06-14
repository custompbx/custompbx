# Use an official Python runtime as a parent image
FROM python:3.9-slim

# Set the working directory in the container
WORKDIR /app

# Install dependencies
RUN apt-get update && apt-get install -y \
    git \
    ffmpeg \
    && rm -rf /var/lib/apt/lists/*

# Clone faster-whisper repository
RUN git clone https://github.com/SYSTRAN/faster-whisper.git

# Install Python dependencies
RUN pip install --no-cache-dir \
    torch \
    CTranslate2 \
    websockets \
    && pip install /app/faster-whisper

# Copy the current directory contents into the container at /app
COPY . /app

# Expose the port the app runs on
EXPOSE 8765

# Run the entrypoint script
CMD ["python", "server.py"]