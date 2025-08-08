package services

import (
	"churras/dto"
	"churras/models"
	"churras/repo"
	"errors"
	"fmt"
)

func GetAllComandas() ([]models.Comanda, error){
	comandas, err := repo.FindAllTabs()
	if err != nil {
		return nil, err
	}
	return comandas, nil
}

func NewTab(comandaRequest dto.ComandaRequest) error{

	//impede de criar comanda sem indentificacao
	if comandaRequest.Identificacao == "" {
		return errors.New("Identificação é obrigatorio para criar uma comanda!")
	} 
	
	err := repo.CreateTab(comandaRequest)

	if err != nil{
		return fmt.Errorf("erro ao criar comanda: %w", err) 
	}
	return nil
}

func AddOrderOnbill(comandaID uint, pedidoID uint) error {
	cp := models.ComandaPedido{
		ComandaID: comandaID,
		PedidoID: pedidoID,
	}

	if cp.ComandaID == 0 || cp.PedidoID == 0 {
		//return erro pro controller
		return errors.New("ComandaID ou PedidoID invalido")
	}

	err :=  repo.AddOrderOnbill(cp)
	if err != nil {
		return errors.New("Erro ao adicionar pedido na comanda" + err.Error())
	}

	return nil
}
