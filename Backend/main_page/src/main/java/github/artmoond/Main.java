package github.artmoond;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.security.servlet.SecurityAutoConfiguration;
import org.springframework.core.io.InputStreamResource;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;

//TIP To <b>Run</b> code, press <shortcut actionId="Run"/> or
// click the <icon src="AllIcons.Actions.Execute"/> icon in the gutter.
@SpringBootApplication(exclude = { SecurityAutoConfiguration.class })// убрать скобочки к продакшену
@RestController
public class Main {
     public static void main(String[] args) {
        SpringApplication.run(Main.class, args);

    }
    @GetMapping("/hello")
     String hello(@RequestParam(value = "name", defaultValue = "World") String name) {
        return String.format("Hello %s!", name);
    }
    @GetMapping(value = "stream/StartAudio", produces= "audio/mpeg")
    public ResponseEntity<InputStreamResource> StartAudio() throws IOException {
        File file = new File(System.getProperty("user.dir") + "/src/main/resources/static/audio/NITSHEMERTW.mp3");
        InputStreamResource resource = new InputStreamResource((new FileInputStream(file)));
        return ResponseEntity.ok().contentLength(file.length()).body(resource);
    }
}
