FROM archlinux:latest

RUN pacman -Syu --noconfirm \
    && pacman -S --noconfirm \
    base-devel \
    bash-completion \
    openssh \
    git \
    go \
    && pacman -Scc --noconfirm

# Enable sudo permission for wheel group
RUN sed -i 's/# %wheel ALL=(ALL) ALL/%wheel ALL=(ALL) ALL/' /etc/sudoers

RUN useradd --create-home --shell /bin/bash --groups wheel archuser
USER archuser

ENTRYPOINT ["/bin/bash"]