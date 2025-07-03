shortlyurl/
├── cmd/                    # Applications principales
│   ├── server/            # Application serveur principal
│   │   └── main.go       # Point d'entrée du serveur
│   ├── cli/              # Interface en ligne de commande (optionnel)
│   │   └── main.go
│   └── worker/           # Worker/daemon background (optionnel)
│       └── main.go
├── internal/              # Code privé de l'application
│   ├── config/           # Configuration de l'application
│   │   └── config.go
│   ├── domain/           # Entités métier et modèles
│   │   ├── user.go
│   │   └── url.go
│   ├── handler/          # Handlers HTTP (controllers)
│   │   ├── auth.go
│   │   ├── url.go
│   │   └── middleware.go
│   ├── service/          # Logique métier
│   │   ├── auth.go
│   │   ├── url.go
│   │   └── analytics.go
│   ├── repository/       # Couche d'accès aux données
│   │   ├── interfaces.go
│   │   ├── postgres/
│   │   │   ├── user.go
│   │   │   └── url.go
│   │   └── redis/
│   │       └── cache.go
│   └── database/         # Connexions et migrations DB
│       ├── connection.go
│       └── migrations/
│           ├── 001_initial.sql
│           └── 002_indexes.sql
├── pkg/                   # Code réutilisable par d'autres projets
│   ├── logger/           # Utilitaire de logging
│   │   └── logger.go
│   ├── jwt/              # Utilitaire JWT
│   │   └── jwt.go
│   ├── validator/        # Validation des données
│   │   └── validator.go
│   └── errors/           # Gestion d'erreurs custom
│       └── errors.go
├── api/                   # Documentation API
│   ├── openapi.yaml      # Spécification OpenAPI/Swagger
│   └── proto/            # Fichiers Protocol Buffer (si gRPC)
├── web/                   # Assets web (optionnel)
│   ├── static/           # Fichiers statiques
│   │   ├── css/
│   │   ├── js/
│   │   └── images/
│   └── templates/        # Templates HTML
├── configs/               # Fichiers de configuration
│   ├── config.yaml       # Configuration par défaut
│   ├── config.prod.yaml  # Configuration production
│   └── .env.example      # Variables d'environnement
├── scripts/               # Scripts de build et déploiement
│   ├── build.sh          # Script de build
│   ├── migrate.sh        # Script de migration DB
│   └── test.sh           # Script de tests
├── deployments/           # Configuration de déploiement
│   ├── docker/           # Fichiers Docker
│   │   └── Dockerfile
│   ├── kubernetes/       # Manifestes Kubernetes
│   │   ├── deployment.yaml
│   │   └── service.yaml
│   ├── docker-compose.yml
│   └── systemd/          # Services systemd
├── test/                  # Tests additionnels
│   ├── integration/      # Tests d'intégration
│   ├── e2e/             # Tests end-to-end
│   └── fixtures/        # Données de test
├── docs/                  # Documentation
│   ├── README.md
│   ├── api.md
│   └── deployment.md
├── vendor/                # Dépendances vendorisées (optionnel)
├── .gitignore
├── go.mod                # Module Go
├── go.sum                # Checksums des dépendances
├── README.md             # Documentation principale
└── Makefile              # Commandes de build/test