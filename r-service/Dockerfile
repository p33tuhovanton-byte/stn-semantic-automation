FROM r-base:4.3.0
WORKDIR /app
COPY install_deps.R .
RUN Rscript install_deps.R
COPY . .
CMD ["Rscript", "main.R"]
