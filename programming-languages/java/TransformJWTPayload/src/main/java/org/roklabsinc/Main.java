package org.roklabsinc;
import java.util.Base64;


// convert payload to POJO
//static class PayloadData {
//    private int user_id;
//    private String email;
//    private String first_name;
//    private String last_name;
//    private String other_name;
//    private String office;
//    private String organization;
//    private String profile_image;
//    private String[] roles;
//    private String[] apps;
//    private int iat;
//    private int exp;
//    private String aud;
//    private String iss;
//    private String sub;
//}
//

public class Main {
    public static void main(String[] args) {

        String tokenString = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMTg1LCJlbWFpbCI6ImVtaWx5LnJvZHJpZ3VlekBhbWFsaXRlY2guY29tIiwiZmlyc3RfbmFtZSI6IkVtaWxpZSIsImxhc3RfbmFtZSI6Ik9wb2t1Iiwib3RoZXJfbmFtZSI6bnVsbCwib2ZmaWNlIjoiMzQiLCJvcmdhbml6YXRpb24iOiI0NTciLCJwcm9maWxlX2ltYWdlIjpudWxsLCJyb2xlcyI6W10sImFwcHMiOltdLCJpYXQiOjE3MTEwMDcxNDYsImV4cCI6MTcxMTA5MzU0NiwiYXVkIjoiaHR0cHM6Ly9hbWFsaXRlY2guY29tLyIsImlzcyI6IkFtYWxpdGVjaCBTU08iLCJzdWIiOiJpbmZvQGFtYWxpdGVjaC5jb20ifQ.O9KdSJgDMKRzQ4W6VLkadUdD1LgOhe6UMTPx2kkg6NzqoKa_ogiteavtuDi0UlXhuD_zHRyysAXTUFjRsHunT6uQfeblrt17b0qUZPLB_HBwAYybIMOc14fD8V-MLyvNjWg_mUmoVeTgvQvjzu897scX7hUNqRHPm6ytahtleG5B_wuuFWa0Qam12ggTnMlZ1KgFT7-fU5_7bk-g5Wc11NDE9Wis6z5rmRS0MKO3-2iaQlcp4ymcIjiDd1AIfTFz4vm79UzO4Cb_heMIly-VMxq-4yFtvkUMukMKCY6qTxO-4hZBGkbmoXBo6mUoXUEWrB9ZPAlVhYSvvdtYiCWtkg"; // vars.get("token")

        // decode token and get payload
        String[] tokenChunks = tokenString.split("\\.");
        Base64.Decoder decoder = Base64.getUrlDecoder();

        String jwtPayload = new String(decoder.decode(tokenChunks[1]));
        System.out.println(jwtPayload);

        // convert the object
        ObjectMapper objectMapper = new ObjectMapper();
        //PayloadData payload = objectMapper.readValue(jwtPayload, PayloadData.class);
    }
}
