package org.roklabsinc.chicago_store;

import org.roklabsinc.Pizza;

public class ChicagoCheezePizza implements Pizza {
    @Override
    public void prepare() {
        System.out.println("Preparing Chicago Cheeze Pizza");
    }

    @Override
    public void bake() {
        System.out.println("Baking Chicago Cheeze Pizza");
    }

    @Override
    public void cut() {
        System.out.println("Cutting Chicago Cheeze Pizza");
    }

    @Override
    public void box() {
        System.out.println("Boxing Chicago Cheeze Pizza");
    }
}
