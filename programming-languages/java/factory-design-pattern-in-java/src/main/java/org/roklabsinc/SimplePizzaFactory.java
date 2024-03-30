package org.roklabsinc;

public class SimplePizzaFactory {
    public Pizza createPizza(PizzaType pizzaType) {
        return switch (pizzaType) {
            case PizzaType.CHEESE -> new CheezePizza();
            case PizzaType.VEGGIE -> new VeggiePizza();
            case PizzaType.PEPPERONI -> new PepperoniPizza();
            case PizzaType.CLAM -> new ClamPizza();
            default -> null;
        };
    }
}
