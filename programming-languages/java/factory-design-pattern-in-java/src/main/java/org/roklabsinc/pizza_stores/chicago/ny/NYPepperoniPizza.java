package org.roklabsinc.ny;

import org.roklabsinc.Pizza;

public class NYPepperoniPizza implements Pizza {
    @Override
    public void prepare() {
        System.out.println("Preparing NY Pepperoni Pizza");
    }

    @Override
    public void bake() {
        System.out.println("Baking NY Pepperoni Pizza");
    }

    @Override
    public void cut() {
        System.out.println("Cut NY Pepperoni Pizza");
    }

    @Override
    public void box() {
        System.out.println("Boxing NY Pepperoni Pizza");
    }
}
