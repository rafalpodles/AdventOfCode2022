package com.example.testcontainers;

import java.io.BufferedReader;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.nio.charset.StandardCharsets;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;

public class Dupa {


  public void main(String[] args)
    try {
      ClassLoader classloader = Thread.currentThread().getContextClassLoader();
      InputStream is = classloader.getResourceAsStream("data");
      assert is != null;
      InputStreamReader streamReader = new InputStreamReader(is, StandardCharsets.UTF_8);
      BufferedReader br = new BufferedReader(streamReader);
      Map<String, Long> map = new HashMap<>();
      String position = "";
      for (String line; (line = br.readLine()) != null; ) {

        if (line.startsWith("$")) {
          if (line.contains("cd")) {
            String[] split = line.split(" ");
            if (Objects.equals(split[2], "..")) {
              int index = position.lastIndexOf('/');
              position = position.substring(0, index);
            } else {
              position = position + "/" + split[2];
              map.put(position, 0L);
            }

          }
        } else {
          if (!line.startsWith("dir")) {
            String[] split = line.split(" ");
            String[] split1 = position.split("/");
            String path = "";
            for (int i = 1; i < split1.length; i++) {
              path = path + "/" + split1[i];
              map.put(path, map.get(path) + Long.parseLong(split[0]));
            }
          }
        }
      }

      long sum = map.values().stream().filter(aLong -> aLong < 100000).mapToLong(aLong -> aLong)
          .sum();
      System.out.println("Sum smaller than 100000: ");
      System.out.println(sum);

      Long totalUsedSpace = map.get("/*");
      Long maxSpace = 70000000L;
      Long requiredSpace = 30000000L;
      Long freeSpace = maxSpace - totalUsedSpace;
      Long spaceToFree = requiredSpace - freeSpace;
      Long spaceToRemove = totalUsedSpace;
      for (Long value : map.values()) {
        if(value>spaceToFree && value<spaceToRemove){
          spaceToRemove = value;
        }
      }
      System.out.println("Space To Remove:");
      System.out.println(spaceToRemove);


    } catch (Exception e) {
      e.printStackTrace();
    }
  }


}
