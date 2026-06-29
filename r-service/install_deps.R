# Скрипт установки зависимостей для R-сервиса
required_packages <- c("plumber", "jsonvalidate", "jsonlite")

install_if_missing <- function(pack) {
  if (!require(pack, character.only = TRUE)) {
    install.packages(pack, repos = "https://r-project.org")
  }
}

invisible(lapply(required_packages, install_if_missing))
print("✅ Все зависимости для R-сервиса успешно установлены!")
