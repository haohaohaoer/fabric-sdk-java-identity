package org.app.network;

import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class ReadTxt {

    public static List<String> ip_port = new ArrayList<>();

    public static void read() {

        String pathname = System.getProperty("user.dir") + "/input.txt";

        try (FileReader reader = new FileReader(pathname);
             BufferedReader br = new BufferedReader(reader)){
            String line;
            while ((line = br.readLine()) != null) {
                 ip_port.add(line);
             }
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
