FROM alpine:latest


RUN apk add --no-cache build-base cmake curl


WORKDIR /app
RUN git clone https://github.com/ggerganov/llama.cpp.git
WORKDIR /app/llama.cpp
RUN make


RUN mkdir -p /app/models
RUN curl -L -o /app/models/tinyllama.bin https://huggingface.co/TheBloke/TinyLlama-1.1B-GGUF/resolve/main/tinyllama-1.1b.Q4_K_M.gguf

EXPOSE 8080

CMD ["./server", "-m", "/app/models/tinyllama.bin", "--host", "0.0.0.0", "--port", "8080"]