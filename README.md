# Loja de camisetas

## Implementações
`Datadog`, `Apache Kafka`, `Swager`

### Users
| Parâmetro    | Tipo        | Descrição            | Valor Padrão      |
|:-------------|:------------|:---------------------|:------------------
| `id`         | `integer`   | Autoincrement        | NULL              |
| `name`       | `string`    | Nome do user         | Nenhum            |
| `email`      | `string`    | Email do user        | Nenhum            |
| `password`   | `string`    | Password do user     | Nenhum            |
| `deleted_at` | `timestamp` | Caso esteja deletado | NULL              |
| `created_at` | `timestamp` | criação              | CURRENT_TIMESTAMP |
| `updated_at` | `timestamp` | ultima atualização   | CURRENT_TIMESTAMP |

### Sizes
| Parâmetro    | Tipo        | Descrição            | Valor Padrão      |
|:-------------|:------------|:---------------------|:------------------
| `id`         | `integer`   | Autoincrement        | NULL              |
| `name`       | `string`    | Nome do tamanho      | Nenhum            |
| `deleted_at` | `timestamp` | Caso esteja deletado | NULL              |
| `created_at` | `timestamp` | criação              | CURRENT_TIMESTAMP |
| `updated_at` | `timestamp` | ultima atualização   | CURRENT_TIMESTAMP |

### Colors
| Parâmetro         | Tipo        | Descrição              | Valor Padrão      |
|:------------------|:------------|:-----------------------|:------------------
| `id`              | `integer`   | Autoincrement          | NULL              |
| `name`            | `string`    | Nome da cor            | Nenhum            |
| `secondary_color` | `string`    | Nome da cor secundaria | Nenhum            |
| `deleted_at`      | `timestamp` | Caso esteja deletado   | NULL              |
| `created_at`      | `timestamp` | criação                | CURRENT_TIMESTAMP |
| `updated_at`      | `timestamp` | ultima atualização     | CURRENT_TIMESTAMP |

### Categories
| Parâmetro     | Tipo        | Descrição              | Valor Padrão      |
|:--------------|:------------|:-----------------------|:------------------
| `id`          | `integer`   | Autoincrement          | NULL              |
| `name`        | `string`    | Nome da categoria      | Nenhum            |
| `slug`        | `string`    | Slug da categoria      | Nenhum            |
| `index`       | `integer`   | Index de ordenação     | Nenhum            |
| `description` | `string`    | Descrição da categoria | NULL              |
| `deleted_at`  | `timestamp` | Caso esteja deletado   | NULL              |
| `created_at`  | `timestamp` | criação                | CURRENT_TIMESTAMP |
| `updated_at`  | `timestamp` | ultima atualização     | CURRENT_TIMESTAMP |

### Sub Categories
| Parâmetro      | Tipo        | Descrição                     | Valor Padrão      |
|:---------------|:------------|:------------------------------|:------------------
| `id`           | `integer`   | Autoincrement                 | NULL              |
| `category_id`  | `integer`   | Referencia para categoria pai | Nenhum            |
| `name`         | `string`    | Nome da categoria             | Nenhum            |
| `slug`         | `string`    | Slug da categoria             | Nenhum            |
| `description`  | `string`    | Descrição da sub-categoria    | NULL              |
| `deleted_at`   | `timestamp` | Caso esteja deletado          | NULL              |
| `created_at`   | `timestamp` | criação                       | CURRENT_TIMESTAMP |
| `updated_at`   | `timestamp` | ultima atualização            | CURRENT_TIMESTAMP |

### Products
| Parâmetro    | Tipo        | Descrição               | Valor Padrão      |
|:-------------|:------------|:------------------------|:------------------
| `id`         | `integer`   | Autoincrement           | NULL              |
| `name`       | `string`    | Nome do produto         | Nenhum            |
| `weight`     | `float`     | Peso do produto         | Nenhum            |
| `width`      | `float`     | Largura do produto      | Nenhum            |
| `height`     | `float`     | Altura do produto       | Nenhum            |
| `price`      | `float`     | Preço do produto        | Nenhum            |
| `deleted_at` | `timestamp` | Caso esteja deletado    | NULL              |
| `created_at` | `timestamp` | criação                 | CURRENT_TIMESTAMP |
| `updated_at` | `timestamp` | ultima atualização      | CURRENT_TIMESTAMP |

### Product Variations
| Parâmetro    | Tipo         | Descrição                      | Valor Padrão      |
|:-------------|:-------------|:-------------------------------|:------------------
| `id`         | `integer`    | Autoincrement                  | NULL              |
| `product_id` | `integer`    | id do produto pai              | Nenhum            |
| `size_id`    | `integer`    | Referencia ao tamanho da peça  | Nenhum            |
| `color_id`   | `integer`    | Referencia ao color caso tenha | NULL              |
| `created_at` | `timestamp`  | criação                        | CURRENT_TIMESTAMP |
| `updated_at` | `timestamp`  | ultima atualização             | CURRENT_TIMESTAMP |

### Product Categories
| Parâmetro         | Tipo        | Descrição               | Valor Padrão      |
|:------------------|:------------|:------------------------|:------------------
| `id`              | `integer`   | Autoincrement           | NULL              |
| `product_id`      | `integer`   | id da variação pai      | Nenhum            |
| `sub_category_id` | `integer`   | id da sub categoria pai | Nenhum            |
| `created_at`      | `timestamp` | criação                 | CURRENT_TIMESTAMP |
| `updated_at`      | `timestamp` | ultima atualização      | CURRENT_TIMESTAMP |

### Product Stocks
| Parâmetro              | Tipo        | Descrição           | Valor Padrão      |
|:-----------------------|:------------|:--------------------|:------------------
| `id`                   | `integer`   | Autoincrement       | NULL              |
| `product_variation_id` | `integer`   | id da variação pai  | Nenhum            |
| `quantity`             | `integer`   | Quantidade em stock | Nenhum            |
| `created_at`           | `timestamp` | criação             | CURRENT_TIMESTAMP |
| `updated_at`           | `timestamp` | ultima atualização  | CURRENT_TIMESTAMP |
