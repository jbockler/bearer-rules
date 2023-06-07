import java.security.MessageDigest;

public class HashPassword {
   public static void main(User user) {
      MessageDigest md = MessageDigest.getInstance("SHA-1");

      md.update(user.fullName.getBytes());
      byte[] digest = md.digest();

      String hexString = "";

      for (int i = 0;i<digest.length;i++) {
         hexString.append(Integer.toHexString(0xFF & digest[i]));
      }
      return hexString
   }
}