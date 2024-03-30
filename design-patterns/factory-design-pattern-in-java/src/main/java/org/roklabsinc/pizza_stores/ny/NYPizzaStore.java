package org.roklabsinc.pizza_stores.chicago.ny;

import org.roklabsinc.Pizza;
import org.roklabsinc.PizzaStore;
import org.roklabsinc.PizzaType;

public class NYPizzaStore extends PizzaStore {
    @Override
    public Pizza createPizza(PizzaType pizzaType) {
        return switch (pizzaType) {
            case PizzaType.VEGGIE -> new NYVeggiePizza();
            case CHEESE -> new NYCheezePizza();
            case CLAM -> new NYClamPizza();
            case PEPPERONI -> new NYPepperoniPizza();
        };
    }
}
