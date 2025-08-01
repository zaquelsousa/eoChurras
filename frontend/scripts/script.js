const url = "/api/comandas";


const allComandas = document.getElementById('comandas');
const conteinercomandasDiv = document.createElement("div");
fetch(url, {
    method: "GET"
}).then((response) => response.json()).then((json) => {

        json.forEach(e => {
            const comandasDiv = document.createElement("div");

            let title = document.createElement("h3");
            btnNewPedido = document.createElement("button");
            btnNewPedido.textContent = "Novo Pedido";
            btnNewPedido.setAttribute('data-id', e.ID)
            btnNewPedido.setAttribute('comanda-name', e.Identificacao);
            btnNewPedido.classList.add("AddPedido");
            title.textContent = "Identificaçao: " + e.Identificacao;
            let preco = document.createElement("p");
            preco.textContent = "Valor: " + e.Preco;
            let status = document.createElement("p");
            status.textContent = "Status: " + e.EstaFechada;
            let date = document.createElement("p");
            
            let data = new Date(e.CreatedAt);
            let d = String(data.getDate()).padStart(2, '0');
            const m = String(data.getMonth() + 1).padStart(2, '0');
            const a = String(data.getFullYear()).slice(-2);
            date.textContent = "Criada em: " + `${d}/${m}/${a}`;
            
            let btnFechar = document.createElement("button")
            btnFechar.classList.add("conteinerBTN")
            btnFechar.textContent = "Fechar Conta" 
            let btnInfo = document.createElement("button");
            btnInfo.classList.add("conteinerBTN", "moreInfo");
            btnInfo.setAttribute("data-id", e.ID);
            btnInfo.textContent = "mais informaçoes"
            let btnContainer = document.createElement("div");
            btnContainer.classList.add("btn-container");
            btnContainer.appendChild(btnFechar);
            btnContainer.appendChild(btnInfo);
            comandasDiv.append(title, btnNewPedido, preco, status, date, btnContainer);
            comandasDiv.classList.add("comandasCard");
            conteinercomandasDiv.appendChild(comandasDiv)
    })

    allComandas.appendChild(conteinercomandasDiv)
})


document.getElementById("comandaForm").addEventListener("submit", function(event){
    event.preventDefault();
    const formData = document.getElementById("Identificacao"); 
        let _comanda = {
        Identificacao: formData.value,
        EstaFechada: false,
        UserID: 1,
        Preco: 0
    }
    
    fetch(url, {
        method: "POST",
        body: JSON.stringify(_comanda),
        headers: {"content-type": "application/json; charset=utf-8"}
    }).then(response => response.text()).then(result => {console.log(result)}).catch(error => {console.error(error)})
})


document.addEventListener("DOMContentLoaded", () => {
    const novaComandaBtn = document.querySelector(".title button");
    const modal = document.getElementById("createComanda");
    const closeButton = document.querySelector(".close-button");
    const createbtn = document.querySelector(".btnCreate");
    novaComandaBtn.addEventListener("click", () => {
        modal.style.display = "block";
    });

    closeButton.addEventListener("click", () => {
        modal.style.display = "none";
    });
    
    createbtn.addEventListener("click", () => {
        modal.style.display = "none";
    })
    window.addEventListener("click", (event) => {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    });
});


document.addEventListener("DOMContentLoaded", () => {
    const socket = new WebSocket(`ws://${location.host}/api/ws`);

    socket.onopen = () => {
        console.log("WebSocket conectado");
    };

    socket.onmessage = (event) => {
        const msg = JSON.parse(event.data);

        switch (msg.tipo) {
            case "comanda":
                renderNovaComanda(msg.dado);
                break;
            case "notificacao":
                showPopup(msg.dado.mensagem);
                break;
            default:
                console.warn("Tipo de mensagem desconhecido:", msg.tipo);
        }
    };

    function showPopup(message) {
        const popup = document.createElement('div');
        const msg = document.createElement("h3");
        msg.innerText = message;
        popup.appendChild(msg);

        popup.style.position = 'fixed';
        popup.style.top = '20px';
        popup.style.left = '50%';
        popup.style.transform = 'translateX(-50%)';
        popup.style.padding = '15px 20px';
        popup.style.backgroundColor = '#333';
        popup.style.color = '#fff';
        popup.style.borderRadius = '8px';
        popup.style.boxShadow = '0 4px 12px rgba(0, 0, 0, 0.2)';
        popup.style.zIndex = '10000';
        popup.style.fontFamily = 'Arial, sans-serif';
        popup.style.fontSize = '14px';
        popup.style.maxWidth = '90%';
        popup.style.textAlign = 'center';
        popup.style.wordWrap = 'break-word';
        popup.style.opacity = '0';
        popup.style.transition = 'opacity 0.3s ease';

        document.body.appendChild(popup);

        requestAnimationFrame(() => {
            popup.style.opacity = '1';
        });

        setTimeout(() => {
            popup.style.opacity = '0';
            setTimeout(() => {
                popup.remove();
            }, 300);
        }, 3000);
    }

    function renderNovaComanda(novaCamanda) {
        const comandasDiv = document.createElement("div");

        let title = document.createElement("h3");
        title.textContent = "Identificaçao: " + novaCamanda.Identificacao;
        let preco = document.createElement("p");
        preco.textContent = "Valor: " + novaCamanda.Preco;
        let status = document.createElement("p");
        status.textContent = "Status: " + novaCamanda.EstaFechada;
        let date = document.createElement("p");
        let data = new Date(novaCamanda.CreatedAt);
        let d = String(data.getDate()).padStart(2, '0');
        const m = String(data.getMonth() + 1).padStart(2, '0');
        const a = String(data.getFullYear()).slice(-2);
        date.textContent = "Criada em: " + `${d}/${m}/${a}`;

        let btnFechar = document.createElement("button");
        btnFechar.classList.add("conteinerBTN");
        btnFechar.textContent = "Fechar Conta";

        let btnInfo = document.createElement("button");
        btnInfo.classList.add("conteinerBTN", "moreInfo");
        btnInfo.setAttribute("data-id", novaCamanda.ID);
        btnInfo.textContent = "mais informaçoes";

        let btnContainer = document.createElement("div");
        btnContainer.classList.add("btn-container");
        btnContainer.appendChild(btnFechar);
        btnContainer.appendChild(btnInfo);

        comandasDiv.append(title, preco, status, date, btnContainer);

        conteinercomandasDiv.appendChild(comandasDiv);
    }
});



const urlProdutos = "/api/produtos";
const urlPedidos = "/api/pedidos";

const btnaddPedido = document.getElementById("addPedido");
const btnAddProduto = document.getElementById("addProduto");


let Allprodutos = []

fetch(urlProdutos, {
    method: "GET"
}).then((r) => r.json()).then((json) => {
    const select = document.getElementById("produtos");
    json.forEach(e => {
        Allprodutos.push(e)
        let opt = document.createElement('option');
        opt.value = e.ID;
        opt.innerHTML = e.Name;
        select.appendChild(opt);
    })
})

let produtos = [];


let comandaID;
let comandaName;
document.addEventListener("DOMContentLoaded", () => {
    const modal = document.getElementById("createPedido");
    const closeButton = document.querySelector(".closeButtonMpedido");
    const createbtn = document.querySelector(".btnCreatePedido");
    const comandaNAME = document.getElementById("comandaName");

    // CARALHO EU ODEIO JS SE FUDER TEM QUE FAZER UMA DELEGAÇAO DE EVENTO PQ OS BTS SAO CIADOS DIN AMICAMENTE PLMDS
    // AI A GENTE VERIFICA NO PAI SE UM FILHO FOI CLICADO SE FUDER LIUNGUAGEM DE MERDA
    // https://www.freecodecamp.org/news/event-delegation-javascript/
    document.body.addEventListener("click", (event) => {
        const target = event.target;
        // AINDA TEM QUE VER SE O CLICK VEIO NA CLASSE FILHA QUE A QUE NOS QUER 
        if (target.classList.contains("AddPedido")) {
            comandaID = parseInt(target.getAttribute("data-id"));
            console.log(typeof(comandaID));
            comandaName = target.getAttribute("comanda-name");
            comandaNAME.value = comandaName;
            console.log(comandaName)
            modal.style.display = "block";
        }
    });


    closeButton.addEventListener("click", () => {
        modal.style.display = "none";
    });

    createbtn.addEventListener("click", () => {
        modal.style.display = "none";
    });

    window.addEventListener("click", (event) => {
        if (event.target === modal) {
            modal.style.display = "none";
        }
    });
});


const ulItensAtuais = document.getElementById("itensAtuais");
let currentClick = 0;
btnAddProduto.addEventListener("click", function(e){
    e.preventDefault();
    let opt = parseInt(document.getElementById("produtos").value)
    let qtd = parseInt(document.getElementById("qtd").value)
    let preco = Allprodutos.find(p => p.ID === opt)
    produtos.push({
        ProdutoID: opt,
        Quantidade: qtd,
        Preco: parseFloat(preco.Preco) * qtd
    })
    let li = document.createElement("li");
    let produtoName = Allprodutos.find(p => p.ID === opt)
    li.textContent = qtd + ": " + produtoName.Name;
    ulItensAtuais.appendChild(li);
    currentClick += 1;
    console.log(produtos)
})

document.getElementById("PedidoForm").addEventListener("submit", function(e){
    e.preventDefault()
    let pedido = {
        StatusPedido: 0,
        ComandaID: comandaID,
        Produtos: produtos
    }

    fetch(urlPedidos, {
        method: "POST",
        body: JSON.stringify(pedido),
        headers: {"content-type": "application/json; charset=utf-8"}
    }).then(response => response.text()).then(result => {console.log(result)}).catch(error => {console.error(error)})

})


