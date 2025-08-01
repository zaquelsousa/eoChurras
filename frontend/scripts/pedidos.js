const url = "/api/comandas";
const urlProdutos = "/api/produtos";
const urlPedidos = "/api/pedidos";


async function loadData(){
    fetch(url, {
        method: "GET"
    }).then((r) => r.json()).then((json) => {
        const select = document.getElementById("comandas");
        json.forEach(e => {
            let opt = document.createElement('option');
            opt.value = e.ID;
            opt.innerHTML = e.Identificacao;
            select.appendChild(opt);
        })
    })


    let todosProdutos = [];

    async function getProdutos(){
        await fetch(urlProdutos, {
            method: "GET"
        }).then((r) => r.json()).then((json) => {
            json.forEach(e => {
                todosProdutos.push(e);
            })
        })
    }

    await getProdutos();

    let todosComandas = [];

    async function getComandas(){
        await fetch(url, {
            method: "GET"
        }).then((r) => r.json()).then((json) => {
            json.forEach(e => {
                todosComandas.push(e);
            })
        })
    }

    await getComandas();

    const pedidos = document.getElementById('Pedidos');
    const conteinerPedidos = document.createElement("div");

    fetch(urlPedidos,  {
        method: "GET"
    }).then((r) => r.json()).then((json) => {
        json.forEach(e => {
            const pedidosDiv = document.createElement("div");
            let comanda = document.createElement("h3");
            let comandaIdentificacao = todosComandas.find(c => c.ID === e.ComandaID)
            comanda.textContent = "Comanda: " + comandaIdentificacao.Identificacao;
            let ulProdutos = document.createElement("ul");
            e.Produtos.forEach(p => {
                let li = document.createElement("li");
                let produtoName = todosProdutos.find(prod => prod.ID === p.ProdutoID);
                li.textContent = "Produdo: " + produtoName.Name + " Quantidade: " + p.Quantidade
                ulProdutos.appendChild(li);
            })
            let btnConcluir = document.createElement("button");
            btnConcluir.classList.add("btnConcluir")
            btnConcluir.setAttribute("data-id", e.ID);
            btnConcluir.textContent = "Feito" 
            let btnCancelar = document.createElement("button");
            btnCancelar.textContent = "Cancelar"
            
            let conteinerBtns = document.createElement("div");
            conteinerBtns.classList.add("conteiner-bts");
            conteinerBtns.append(btnConcluir, btnCancelar);
            pedidosDiv.append(comanda, ulProdutos, conteinerBtns);
            conteinerPedidos.appendChild(pedidosDiv);
        })
        pedidos.appendChild(conteinerPedidos);
    })

    const socket = new WebSocket(`ws://${location.host}/api/ws`);

    socket.onopen = () => {
        console.log("concectado parsa");
    };

    socket.onmessage = (event) => {
        const msg = JSON.parse(event.data);

        if (msg.tipo === "pedido") {
            renderNovoPedido(msg.dado);
        }
    };

    function renderNovoPedido(novoPedido) {
        const pedidosDiv = document.createElement("div");

        let comanda = document.createElement("h3");
        comanda.textContent = "Comanda: " + novoPedido.ComandaID;

        let ulProdutos = document.createElement("ul");
        novoPedido.Produtos.forEach(p => {
            let li = document.createElement("li");
            let produtoName = todosProdutos.find(prod => prod.ID === p.ProdutoID);
            li.textContent = "Produdo: " + produtoName.Name + " Quantidade: " + p.Quantidade;
            ulProdutos.appendChild(li);
        });

        let btnConcluir = document.createElement("button");
        btnConcluir.textContent = "Feito";
        btnConcluir.classList.add("btnConcluir");
        btnConcluir.setAttribute("data-id", novoPedido.ID);

        let btnCancelar = document.createElement("button");
        btnCancelar.textContent = "Cancelar";

        let conteinerBtns = document.createElement("div");
        conteinerBtns.classList.add("conteiner-bts");
        conteinerBtns.append(btnConcluir, btnCancelar);

        pedidosDiv.append(comanda, ulProdutos, conteinerBtns);

        conteinerPedidos.appendChild(pedidosDiv);
    }

    const pronto = "/api/pedidos/"
    conteinerPedidos.addEventListener("click", function(e) {
        if (e.target && e.target.classList.contains("btnConcluir")) {
            const id = e.target.getAttribute("data-id");
            console.log("Pedido concluído:", id);

            fetch(`${pronto}${id}/pronto`, {
                method: "GET"
            }).then(r => r.text()).then(console.log);
        }
    });

}

loadData()
