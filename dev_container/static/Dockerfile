# Use an official base image
FROM ubuntu:22.04

# Set environment variables to avoid interactive prompts during installation
ENV DEBIAN_FRONTEND=noninteractive

# Update and install base tools
RUN apt-get update && apt-get install -y \
    software-properties-common \
    curl wget zip unzip git vim nano \
    net-tools iputils-ping dnsutils \
    build-essential \
    python3 python3-pip python3-venv \
    gnupg2 jq htop tree \
    && apt-get clean

# Install common tools
RUN curl -sL https://aka.ms/InstallAzureCLIDeb | bash && \
    curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip" && \
    unzip awscliv2.zip && ./aws/install && rm -rf awscliv2.zip ./aws && \
    curl -LO "https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl" && \
    chmod +x kubectl && mv kubectl /usr/local/bin/ && \
    curl -LO "https://github.com/ahmetb/kubectx/releases/download/v0.9.4/kubens" && \
    curl -LO "https://github.com/ahmetb/kubectx/releases/download/v0.9.4/kubectx" && \
    chmod +x kubens kubectx && mv kubens kubectx /usr/local/bin/ && \
    curl -LO https://github.com/derailed/k9s/releases/download/v0.27.4/k9s_Linux_amd64.tar.gz && \
    tar -zxvf k9s_Linux_amd64.tar.gz && mv k9s /usr/local/bin/ && \
    curl -LO https://releases.hashicorp.com/terraform/1.6.0/terraform_1.6.0_linux_amd64.zip && \
    unzip terraform_1.6.0_linux_amd64.zip && mv terraform /usr/local/bin/ && rm terraform_1.6.0_linux_amd64.zip

# Install Go
RUN curl -LO https://golang.org/dl/go1.20.8.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.20.8.linux-amd64.tar.gz && \
    rm go1.20.8.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

# Install Skaffold
RUN curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
    chmod +x skaffold && \
    mv skaffold /usr/local/bin/


# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs

# Install PowerShell
RUN wget -q https://packages.microsoft.com/config/ubuntu/22.04/packages-microsoft-prod.deb && \
    dpkg -i packages-microsoft-prod.deb && \
    apt-get update && apt-get install -y powershell && \
    rm packages-microsoft-prod.deb

# Install Ansible
RUN apt-get install -y ansible

# Add additional common DevOps tools
RUN apt-get install -y \
    docker.io \
    docker-compose \
    make \
    sshpass \
    openjdk-11-jdk

# Configure Vim
COPY vimrc /root/.vimrc

# Add additional useful tools as needed
RUN pip install --upgrade pip && pip install httpie

# Add custom CLI scripts
COPY welcome.sh /usr/local/bin/welcome.sh
COPY list-tools /usr/local/bin/list-tools

# Make scripts executable
RUN chmod +x /usr/local/bin/welcome.sh /usr/local/bin/list-tools

# Clean up
RUN apt-get clean && rm -rf /var/lib/apt/lists/*

# Set the entrypoint to the welcome script
ENTRYPOINT ["/usr/local/bin/welcome.sh"]
