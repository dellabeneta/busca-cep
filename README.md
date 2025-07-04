## Consulta CEP

![Pipeline](https://github.com/dellabeneta/busca-cep/actions/workflows/main.yaml/badge.svg)
![GHCR version](https://img.shields.io/badge/ghcr.io%2Fdellabeneta%2Fcep-v7-blue)

Um serviço web para consulta de CEP brasileiro, desenvolvido em Go. Consulte rapidamente informações de endereço a partir do CEP usando uma interface web simples e responsiva.

### Aviso Importante

**Este serviço é apenas para fins educacionais e de teste.** Não utilize para fins ilegais. O projeto foi criado para aprendizado e demonstração de tecnologias.

### Funcionalidades

- Consulta de CEPs brasileiros via API pública (ViaCEP)
- Interface web responsiva para desktop e mobile
- Exibição de logradouro, bairro, cidade, UF e DDD
- Backend em Go com endpoint `/consultar`
- Pronto para rodar em Docker e Kubernetes
- CI/CD com GitHub Actions

### Começando

**Pré-requisitos**
- Go 1.24.3 ou superior
- Docker (opcional)
- Kubernetes/k3s (opcional)

**Instalação Local**
```bash
# Clone o repositório
git clone https://github.com/seu-dellabeneta/busca-cep.git
cd busca-cep

# Execute a aplicação
go run main.go
```

A aplicação estará disponível em `http://localhost:8080`

**Usando Docker**
```bash
# Construa a imagem
docker build -t busca-cep .

# Execute o container
docker run -p 8080:8080 busca-cep
```

**Deploy no Kubernetes**
```bash
# Crie o namespace
kubectl apply -f k3s/namespace.yaml

# Deploy da aplicação
kubectl apply -f k3s/deployment.yaml
kubectl apply -f k3s/service.yaml
```

A aplicação estará disponível na porta `30083` do cluster.

### Tecnologias

- **Backend**: Go 1.24.3
- **Frontend**: HTML5, CSS3
- **Container**: Docker Alpine
- **Orquestração**: Kubernetes/k3s
- **API Externa**: ViaCEP

### Como Funciona

O serviço implementa a seguinte lógica:

1. **Recebe o CEP**: O usuário digita o CEP na interface web.
2. **Validação**: O backend valida o formato do CEP.
3. **Consulta Externa**: O backend consulta a API ViaCEP.
4. **Resposta**: Os dados do endereço são retornados em JSON e exibidos na interface.

### Configuração

O serviço pode ser configurado através das seguintes variáveis de ambiente:

| Variável | Descrição | Padrão |
|----------|-----------|--------|
| PORT     | Porta do servidor | 8080 |

### Estrutura do Projeto
```
della@ubuntu:~/projetos/busca-cep$ tree
.
├── Dockerfile
├── go.mod
├── k3s
│   ├── deployment.yaml
│   ├── namespace.yaml
│   └── service.yaml
├── main.go
├── nuke.sh
├── README.md
└── static
    ├── index.html
    └── style.css

3 directories, 10 files

```

### Scripts Úteis

**nuke.sh**: Script para limpeza completa do Docker (containers, imagens, volumes e redes)

```bash
chmod +x nuke.sh
./nuke.sh
```

### Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais