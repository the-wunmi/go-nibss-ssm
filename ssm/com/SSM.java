import org.graalvm.nativeimage.IsolateThread;
import org.graalvm.nativeimage.c.function.CEntryPoint;
import org.graalvm.nativeimage.c.type.CCharPointer;
import org.graalvm.nativeimage.c.type.CTypeConversion;
import nfp.ssm.core.SSMLib;
import java.security.Provider;
import java.security.Security;
import java.security.KeyPairGenerator;
import java.security.NoSuchProviderException;
import java.security.NoSuchAlgorithmException;

public class SSM {
    @CEntryPoint(name = "generateKeyPair")
    private static void generateKeyPair(IsolateThread thread, CCharPointer publicKeyLocation, CCharPointer privateKeyLocation, CCharPointer userId, CCharPointer password) {
        final SSMLib ssmLib = new SSMLib(CTypeConversion.toJavaString(publicKeyLocation), CTypeConversion.toJavaString(privateKeyLocation));
        ssmLib.generateKeyPair(CTypeConversion.toJavaString(userId), CTypeConversion.toJavaString(password));
    }

    @CEntryPoint(name = "encryptMessage")
    private static CCharPointer encryptMessage(IsolateThread thread, CCharPointer publicKeyLocation, CCharPointer message) {
        final SSMLib ssmLib = new SSMLib(CTypeConversion.toJavaString(publicKeyLocation), "");
        final String encrypted  = ssmLib.encryptMessage(CTypeConversion.toJavaString(message));
        CCharPointer encryptedPtr = (CTypeConversion.toCString(encrypted)).get();
        return encryptedPtr;
    }

    @CEntryPoint(name = "decryptMessage")
    private static CCharPointer decryptMessage(IsolateThread thread, CCharPointer privateKeyLocation, CCharPointer password, CCharPointer message) {
        final SSMLib ssmLib = new SSMLib("", CTypeConversion.toJavaString(privateKeyLocation));
        final String decrpyted  = ssmLib.decryptFile(CTypeConversion.toJavaString(message), CTypeConversion.toJavaString(password));
        CCharPointer decrpytedPtr = (CTypeConversion.toCString(decrpyted)).get();
        return decrpytedPtr;
    }

    public static void main(String[] args){}
}