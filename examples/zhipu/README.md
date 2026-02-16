# Zhipu GLM Example

Este exemplo demonstra como usar os modelos GLM da Zhipu (Z.ai) através da biblioteca gollm.

## Pré-requisitos

1. Obtenha uma API key da Z.ai em https://platform.z.ai
2. Configure a variável de ambiente:

```bash
export ZHIPU_API_KEY=sua-chave-api-aqui
```

## Como Executar

```bash
cd examples/zhipu
go run main.go
```

## O que este exemplo demonstra

1. **Geração de texto simples**: Enviar uma pergunta e receber uma resposta
2. **Uso de system prompt**: Configurar um prompt de sistema para definir o comportamento do modelo
3. **Conversa estruturada**: Usar prompts com diretivas específicas para controlar a formatação da resposta

## Modelos Disponíveis

- `glm-4.7`: Modelo mais recente da Zhipu
- `glm-4`: Versão anterior do GLM-4
- Outros modelos GLM disponíveis na plataforma Z.ai

## Mais Informações

- Documentação da Z.ai: https://platform.z.ai
- Documentação da gollm: https://github.com/teilomillet/gollm
