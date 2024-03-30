package org.roklabsinc.pizza_stores.chicago;

import org.roklabsinc.Pizza;

public class ChicagoPepperoniPizza implements Pizza {
    @Override
    public void prepare() {
        System.out.println("Preparing Chicago Pepperoni Pizza");
    }

    @Override
    public void bake() {
        System.out.println("Baking Chicago Pepperoni Pizza");
    }

    @Override
    public void cut() {
        System.out.println("Cutting Chicago Pepperoni Pizza");
    }

    @Override
    public void box() {
        System.out.println("Boxing Chicago Pepperoni Pizza");
    }
}
