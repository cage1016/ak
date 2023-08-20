FROM gcr.io/buildpacks/gcp/run
USER root
RUN apt-get update && apt-get install -y --no-install-recommends \
  zip && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*
USER cnb