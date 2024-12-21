package main

import "fmt"

type Paiement interface {
	EffectuerPaiement(montant float64) string
}

type CarteDeCredit struct {
	NumeroCarte string
}

func (c CarteDeCredit) EffectuerPaiement(montant float64) string {
	return fmt.Sprintf("Paiement de %.2f€ effectué avec succès par carte de crédit (Carte: %s).", montant, c.NumeroCarte)
}

type PayPal struct {
	Email string
}

func (p PayPal) EffectuerPaiement(montant float64) string {
	return fmt.Sprintf("Paiement de %.2f€ réalisé avec succès via PayPal (Email: %s).", montant, p.Email)
}

type CryptoMonnaie struct {
	AdresseWallet string
}

func (crypto CryptoMonnaie) EffectuerPaiement(montant float64) string {
	return fmt.Sprintf("Paiement de %.2f€ effectué via crypto-monnaie (Wallet: %s).", montant, crypto.AdresseWallet)
}

func TraiterPaiement(methodePaiement Paiement, montant float64) {
	resultat := methodePaiement.EffectuerPaiement(montant)
	fmt.Println(resultat)
}

func main() {
	carte := CarteDeCredit{NumeroCarte: "1234-5678-9012-3456"}
	paypal := PayPal{Email: "utilisateur@example.com"}
	crypto := CryptoMonnaie{AdresseWallet: "0xABCDEF1234567890"}

	montant := 150.75

	fmt.Println("Test des paiements :")

	TraiterPaiement(carte, montant)
	TraiterPaiement(paypal, montant)
	TraiterPaiement(crypto, montant)
}
