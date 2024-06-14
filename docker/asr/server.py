import asyncio
import websockets
import json
import numpy as np
import logging
from faster_whisper import WhisperModel

# Configure logging
logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

# Define model path
model_name = "base.en"

# Load the Whisper model
try:
    model = WhisperModel(
        model_name,
        device="cpu",  # Use "cuda" if you have a GPU available
        compute_type="float32"  # You can also use "int8"
    )
    logger.info("Model loaded successfully.")
except Exception as e:
    logger.error(f"Error loading model: {e}")
    raise

# Buffer to store incoming audio data
audio_buffer = np.array([], dtype=np.int16)

async def process_audio(websocket, path):
    global audio_buffer
    logger.info(f"Client connected: {websocket.remote_address}")
    try:
        async for message in websocket:
            logger.debug("Received audio data")

            try:
                # Validate that message contains L16 binary audio data
                if isinstance(message, (bytes, bytearray)):
                    audio_data = np.frombuffer(message, dtype=np.int16)
                    logger.debug(f"Validated audio data length: {audio_data.size}")
                else:
                    logger.error("Received data is not in binary L16 format")
                    continue

                # Append received audio data to the buffer
                audio_buffer = np.append(audio_buffer, audio_data)
                logger.debug(f"Buffered audio data length: {audio_buffer.size}")

                # Process the audio data if buffer length exceeds a threshold (e.g., 1 second)
                if audio_buffer.size >= 16000:  # Assuming 16kHz sample rate
                    logger.info("Processing buffered audio data")
                    audio_to_process = audio_buffer.copy()
                    audio_buffer = np.array([], dtype=np.int16)  # Clear buffer

                    # Convert the buffer to float32 for Whisper processing
                    audio_to_process = audio_to_process.astype(np.float32) / 32768.0

                    segments, _ = model.transcribe(audio_to_process)
                    logger.info("Buffered audio data transcribed")

                    # Extract recognized text
                    recognized_text = " ".join([segment.text for segment in segments])
                    logger.info(f"Recognized text: {recognized_text}")

                    # Send back the recognized text as JSON
                    response = json.dumps({"recognized_text": recognized_text})
                    await websocket.send(response)
                    logger.info("Sent recognized text back to client")
            except Exception as e:
                logger.error(f"Error processing audio data: {str(e)}")
    except websockets.ConnectionClosed as e:
        logger.info(f"Client disconnected: {websocket.remote_address}")
    except Exception as e:
        logger.error(f"Error: {str(e)}")
    finally:
        await websocket.close()

# Start the WebSocket server
start_server = websockets.serve(process_audio, "0.0.0.0", 8765)

logger.info("WebSocket server started")
asyncio.get_event_loop().run_until_complete(start_server)
asyncio.get_event_loop().run_forever()
