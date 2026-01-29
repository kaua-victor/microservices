# 🚀 Projeto de Microsserviços com gRPC (Order, Payment e Shipping)

Este projeto implementa uma arquitetura de **microsserviços em Go**, utilizando **gRPC**, **Docker** e **Docker Compose**, seguindo princípios de **arquitetura hexagonal**.

Os microsserviços implementados são:

* **Order**: responsável por receber pedidos, validar itens e acionar o envio
* **Payment**: responsável por processar pagamentos
* **Shipping**: responsável por calcular o prazo de entrega

---

## 🧱 Arquitetura Geral

Fluxo principal:

1. Cliente envia um pedido para o **Order**
2. Order valida se os produtos existem
3. Order valida quantidade máxima de itens (≤ 50)
4. Order chama o **Payment** para efetuar o pagamento
5. Se o pagamento for bem-sucedido, Order chama o **Shipping**
6. Shipping calcula o prazo de entrega
7. Order finaliza o pedido com status `Paid` ou `Canceled`

Comunicação entre serviços é feita via **gRPC**.

---

## 🐳 Pré-requisitos

Antes de executar o projeto, certifique-se de ter instalado:

* Docker
* Docker Compose

Verifique com:

```bash
docker --version
docker compose version
```

---

## ▶️ Como executar o projeto

### 1️⃣ Subir todos os serviços

Na raiz do projeto (onde está o `docker-compose.yaml`), execute:
bash
```
docker compose up --build -d
```

Este comando irá:

* Criar os containers
* Subir o banco MySQL
* Executar os microsserviços Order, Payment e Shipping

---

## 🌐 Portas utilizadas

| Serviço  | Porta |
| -------- | ----- |
| Order    | 3000  |
| Payment  | 3001  |
| Shipping | 3002  |
| MySQL    | 3306  |

---

## 🗄️ Banco de Dados

O banco MySQL é iniciado automaticamente via Docker.

As bases criadas são:

* `order`
* `payment`
* `shipping`

Além disso, existe a tabela `products`, usada pelo microsserviço Order para validação de itens:

```sql
CREATE TABLE products (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    product_code VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO products (product_code) VALUES
('P1'), ('P2'), ('P3');
```

---

## 🧪 Como testar os serviços

Baixe um grpcurl de acordo com a sua máquina no link a seguir: https://github.com/fullstorydev/grpcurl/releases, e colocar na pasta de arquivos de programas do disco local c
OBS: nesse projeto, foi usado o para windows 1.9.3, dependendo da sua versão, terá que alterar o caminho em alguns códigos do teste

### 🔹 Criar pedidos
powershell na pasta de microservices

& "C:\Program Files\grpcurl_1.9.3_windows_x86_64\grpcurl.exe" -d '{\"costumer_id\":123,\"order_items\":[{\"product_code\":\"P1\",\"quantity\":4,\"unit_price\":12}]}' -plaintext localhost:3000 Order/Create

& "C:\Program Files\grpcurl_1.9.3_windows_x86_64\grpcurl.exe" `
-d '{\"costumer_id\":123,\"order_items\":[
    {\"product_code\":\"P1\",\"quantity\":30,\"unit_price\":10},
    {\"product_code\":\"P2\",\"quantity\":15,\"unit_price\":5}
]}' `
-plaintext localhost:3000 Order/Create

### 🔹 Teste de erro – produto inexistente

& "C:\Program Files\grpcurl_1.9.3_windows_x86_64\grpcurl.exe" -d '{\"costumer_id\":123,\"order_items\":[{\"product_code\":\"P\",\"quantity\":4,\"unit_price\":12}]}' -plaintext localhost:3000 Order/Create

### 🔹 Teste de erro – valor acima de 1000

& "C:\Program Files\grpcurl_1.9.3_windows_x86_64\grpcurl.exe" `
-d '{\"costumer_id\":123,\"order_items\":[
    {\"product_code\":\"P1\",\"quantity\":30,\"unit_price\":10},
    {\"product_code\":\"P2\",\"quantity\":15,\"unit_price\":500}
]}' `
-plaintext localhost:3000 Order/Create

### 🔹 Teste de erro – mais de 50 pedidos

& "C:\Program Files\grpcurl_1.9.3_windows_x86_64\grpcurl.exe" `
-d '{\"costumer_id\":123,\"order_items\":[
    {\"product_code\":\"P1\",\"quantity\":30,\"unit_price\":10},
    {\"product_code\":\"P2\",\"quantity\":45,\"unit_price\":10}
]}' `
-plaintext localhost:3000 Order/Create

## ⚠️ Regras implementadas

* Máximo de **50 itens por pedido**
* Pagamento com timeout e retry
* Shipping só é chamado após pagamento bem-sucedido
* Validação de existência de produtos no banco
* Status do pedido:

  * `Pending`
  * `Paid`
  * `Canceled`

---

## 📦 Estrutura do Projeto

Cada microsserviço segue arquitetura hexagonal:

```
cmd/
internal/
  adapters/
  application/
  ports/
```

---

## 📌 Observações Finais

* A comunicação entre microsserviços é feita exclusivamente via gRPC
* O projeto está preparado para execução local ou via Docker

---

## ✅ Conclusão

Este projeto demonstra conceitos fundamentais de microsserviços:

* Comunicação resiliente
* Separação de responsabilidades
* Arquitetura hexagonal
* Deploy com Docker

---

📚 Projeto desenvolvido para fins acadêmicos.
