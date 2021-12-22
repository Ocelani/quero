# Descritivo do Teste

1. Criar um CRUD de Empresa, com os seguintes campos:
- Id (Auto Incremento)
- Nome (Limite de 200 caracteres)
- Endereço completo (capturado através do CEP, via API https://viacep.com.br/) onde ao digitar o CEP, o restante do endereço vem da API, com inserção extra somente de número e complemento do endereço.
- Telefone, com máscara e validação que pode aceitar a quantidade de dígitos de celulares e telefones fixos, com DDD.

2. Criar um CRUD de Funcionários, com:
- Nome (até 200 caracteres)
- Cargo (Lista fixa com pelo menos 3 opções. Ex: Programador, Designer, Administração) 
- Salário (Entrada de dados de valores, formato: R$ 1.000,00)

### Objetivos (obrigatórios):
- Ao visualizar as propriedades de uma empresa, deve ser mostrada a lista de funcionários da mesma.
- Ao clicar no funcionário, deve se mostrar os dados do mesmo, e o nome da empresa que ele faz parte.
- Tanto o cadastro de novo funcionário quanto a lista de funcionários, deve estar dentro dos detalhes da empresa.
- Um funcionário de uma empresa não deve ser visto na lista de outra empresa.
- O projeto deve ser publicado em algum host (qualquer um) para ser acessado online.

#### Bônus (não obrigatório mas será um diferencial):
- Projeto comitado no Github ou no Bitbucket.
- Área de login para acesso individual para visualização de cada empresa. Ex: Login 1, vë dados da empresa 1. Login 2, vê dados da empresa 2.
- Testes unitários do projeto.
- Paginação e busca nas listas.

### Framework/linguagem Back End
Qualquer uma da sua preferência. Trabalhamos aqui com C#, PHP NodeJS e Ruby, mas pode ser o que se sentir mais confortável.

### Front End 
Simplesmente para mostrar na tela. Não se preocupe com a “beleza”, mas sim com a funcionalidade. Templates, frameworks ou qualquer coisa que facilite o trabalho são bem vindos.
Prazo para disponibilização: 5 dias úteis a partir da conversa inicial.
A não entrega ou falta de retorno, será considerada desistência.Boa sorte!