package org.roklabsinc.pizza_stores.chicago;

import org.roklabsinc.Pizza;

public class ChicagoVeggiePizza implements Pizza {
    @Override
    public void prepare() {
        System.out.println("Preparing Chicago Veggie Pizza");
    }

    @Override
    public void bake() {
        System.out.println("Baking Chicago Veggie Pizza");
    }

    @Override
    public void cut() {
        System.out.println("Cutting Chicago Veggie Pizza");
    }

    @Override
    public void box() {
        System.out.println("Boxing Chicago Veggie Pizza");
    }
}
