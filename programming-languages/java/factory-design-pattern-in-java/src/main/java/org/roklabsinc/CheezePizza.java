package org.roklabsinc;

public class CheezePizza implements Pizza {
    @Override
    public void prepare() {
        System.out.println("Preparing Cheeze Pizza");
    }

    @Override
    public void bake() {
        System.out.println("Baking Cheeze Pizza");
    }

    @Override
    public void cut() {
        System.out.println("Cutting Cheeze Pizza");
    }

    @Override
    public void box() {
        System.out.println("Boxing Cheeze Pizza for Delivery");
    }
}
