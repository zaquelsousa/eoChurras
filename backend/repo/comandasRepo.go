package repo

import (
	"churras/database"
	"churras/dto"
	"churras/models"

)



func FindAllTabs() ([]models.Comanda, error){
	var comandas []models.Comanda
	db := database.GetDB()

	result := db.Find(&comandas)
	return comandas, result.Error
}

func CreateTab(comandaReq dto.ComandaRequest) error{
	db := database.GetDB()
	comanda := models.Comanda{
		Identificacao: comandaReq.Identificacao,
		EstaFechada: comandaReq.EstaFechada,
		UserID: comandaReq.UserID,
		Valor: comandaReq.Valor,
	}

	if err := db.Create(&comanda).Error; err != nil {
		return err
	}

	for _, c := range comandaReq.Pedidos {
		cp := models.ComandaPedido{
			ComandaID: comanda.ID,
			PedidoID: c.PedidoID,
		}

		if err := db.Create(&cp).Error; err != nil {
			return err
		}
	}

	return nil
}


func AddOrderOnbill(cp models.ComandaPedido) error {
	db := database.GetDB()

	if err := db.Create(cp).Error; err != nil{
		return  err
	}

	return nil
}
