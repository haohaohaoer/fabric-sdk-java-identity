## Java SDK Code
### org.app.chaincode.invocation
> InvokeChaincode
    
    package org.app.chaincode.invocation;
    
    public class InvokeChaincode {
    
    }
> InvokeChaincode0
    
    package org.app.chaincode.invocation;
    
    import org.json.JSONObject;
    import org.app.client.CAClient;
    import org.app.client.ChannelClient;
    import org.app.client.FabricClient;
    import org.app.config.Config;
    import org.app.user.UserContext;
    import org.app.util.Util;
    import org.hyperledger.fabric.sdk.*;
    import java.util.Collection;
    import java.util.HashMap;
    import java.util.Map;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    import static java.nio.charset.StandardCharsets.UTF_8;
    
    public class InvokeChaincode0 {
    
        private static final byte[] EXPECTED_EVENT_DATA = "!".getBytes(UTF_8);
        private static final String EXPECTED_EVENT_NAME = "event";
        private static FabricClient fabClient_invoke;
        private static ChannelClient channelClient_invoke;
        static boolean init_judge = true;
    
        public void invokeInit(){
            try {
                Util.cleanUp();
                String caUrl = Config.CA_ORG1_URL;
                CAClient caClient = new CAClient(caUrl, null);
                UserContext adminUserContext = new UserContext();
                adminUserContext.setName(Config.ADMIN);
                adminUserContext.setAffiliation(Config.ORG1);
                adminUserContext.setMspId(Config.ORG1_MSP);
                caClient.setAdminUserContext(adminUserContext);
                adminUserContext = caClient.enrollAdminUser(Config.ADMIN, Config.ADMIN_PASSWORD);
                fabClient_invoke = new FabricClient(adminUserContext);
                channelClient_invoke = fabClient_invoke.createChannelClient(Config.CHANNEL_NAME);
                Channel channel = channelClient_invoke.getChannel();
                Peer peer = fabClient_invoke.getInstance().newPeer("peer0.org1.example.com", "grpc://localhost:7051");
                Orderer orderer = fabClient_invoke.getInstance().newOrderer(Config.ORDERER_NAME, Config.ORDERER_URL);
                channel.addPeer(peer);
                channel.addOrderer(orderer);
                channel.initialize();
                init_judge = false;
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
        public void invoke(String fcn, JSONObject jsonObject){
            try {
                if(init_judge) {
                    invokeInit();
                }
                TransactionProposalRequest request = fabClient_invoke.getInstance().newTransactionProposalRequest();
                ChaincodeID ccid = ChaincodeID.newBuilder().setName(Config.CHAINCODE_1_NAME).build();
                request.setChaincodeID(ccid);
                request.setFcn(fcn);
                String[] arguments = new String[4];
                arguments[0] = jsonObject.getString("identifier");
                arguments[1] = jsonObject.getString("syn_data");
                arguments[2] = jsonObject.getString("id_type");
                arguments[3] = jsonObject.getString("root_id");
                request.setArgs(arguments);
                request.setProposalWaitTime(1000);
                Map<String, byte[]> tm2 = new HashMap<>();
                tm2.put("HyperLedgerFabric", "TransactionProposalRequest:JavaSDK".getBytes(UTF_8));
                tm2.put("method", "TransactionProposalRequest".getBytes(UTF_8));
                tm2.put("result", ":)".getBytes(UTF_8));
                tm2.put(EXPECTED_EVENT_NAME, EXPECTED_EVENT_DATA);
                request.setTransientMap(tm2);
                Collection<ProposalResponse> responses = channelClient_invoke.sendTransactionProposal(request);
                for (ProposalResponse res: responses) {
                    ChaincodeResponse.Status status = res.getStatus();
                    Logger.getLogger(InvokeChaincode.class.getName()).log(Level.INFO,"Invoked initIdentity on "+Config.CHAINCODE_1_NAME + ". Status - " + status);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
> InvokeChaincode1
    
    package org.app.chaincode.invocation;
    
    import org.json.JSONObject;
    import org.app.client.CAClient;
    import org.app.client.ChannelClient;
    import org.app.client.FabricClient;
    import org.app.config.Config;
    import org.app.user.UserContext;
    import org.app.util.Util;
    import org.hyperledger.fabric.sdk.*;
    import java.util.Collection;
    import java.util.HashMap;
    import java.util.Map;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    import static java.nio.charset.StandardCharsets.UTF_8;
    
    public class InvokeChaincode0 {
    
        private static final byte[] EXPECTED_EVENT_DATA = "!".getBytes(UTF_8);
        private static final String EXPECTED_EVENT_NAME = "event";
        private static FabricClient fabClient_invoke;
        private static ChannelClient channelClient_invoke;
        static boolean init_judge = true;
    
        public void invokeInit(){
            try {
                Util.cleanUp();
                String caUrl = Config.CA_ORG1_URL;
                CAClient caClient = new CAClient(caUrl, null);
                UserContext adminUserContext = new UserContext();
                adminUserContext.setName(Config.ADMIN);
                adminUserContext.setAffiliation(Config.ORG1);
                adminUserContext.setMspId(Config.ORG1_MSP);
                caClient.setAdminUserContext(adminUserContext);
                adminUserContext = caClient.enrollAdminUser(Config.ADMIN, Config.ADMIN_PASSWORD);
                fabClient_invoke = new FabricClient(adminUserContext);
                channelClient_invoke = fabClient_invoke.createChannelClient(Config.CHANNEL_NAME);
                Channel channel = channelClient_invoke.getChannel();
                Peer peer = fabClient_invoke.getInstance().newPeer("peer1.org1.example.com", "grpc://localhost:7051");
                Orderer orderer = fabClient_invoke.getInstance().newOrderer(Config.ORDERER_NAME, Config.ORDERER_URL);
                channel.addPeer(peer);
                channel.addOrderer(orderer);
                channel.initialize();
                init_judge = false;
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
        public void invoke(String fcn, JSONObject jsonObject){
            try {
                if(init_judge) {
                    invokeInit();
                }
                TransactionProposalRequest request = fabClient_invoke.getInstance().newTransactionProposalRequest();
                ChaincodeID ccid = ChaincodeID.newBuilder().setName(Config.CHAINCODE_1_NAME).build();
                request.setChaincodeID(ccid);
                request.setFcn(fcn);
                String[] arguments = new String[4];
                arguments[0] = jsonObject.getString("identifier");
                arguments[1] = jsonObject.getString("syn_data");
                arguments[2] = jsonObject.getString("id_type");
                arguments[3] = jsonObject.getString("root_id");
                request.setArgs(arguments);
                request.setProposalWaitTime(1000);
                Map<String, byte[]> tm2 = new HashMap<>();
                tm2.put("HyperLedgerFabric", "TransactionProposalRequest:JavaSDK".getBytes(UTF_8));
                tm2.put("method", "TransactionProposalRequest".getBytes(UTF_8));
                tm2.put("result", ":)".getBytes(UTF_8));
                tm2.put(EXPECTED_EVENT_NAME, EXPECTED_EVENT_DATA);
                request.setTransientMap(tm2);
                Collection<ProposalResponse> responses = channelClient_invoke.sendTransactionProposal(request);
                for (ProposalResponse res: responses) {
                    ChaincodeResponse.Status status = res.getStatus();
                    Logger.getLogger(InvokeChaincode.class.getName()).log(Level.INFO,"Invoked initIdentity on "+Config.CHAINCODE_1_NAME + ". Status - " + status);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
> InvokeChaincode2
 
    package org.app.chaincode.invocation;
    
    import org.json.JSONObject;
    import org.app.client.CAClient;
    import org.app.client.ChannelClient;
    import org.app.client.FabricClient;
    import org.app.config.Config;
    import org.app.user.UserContext;
    import org.app.util.Util;
    import org.hyperledger.fabric.sdk.*;
    import java.util.Collection;
    import java.util.HashMap;
    import java.util.Map;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    import static java.nio.charset.StandardCharsets.UTF_8;
    
    public class InvokeChaincode0 {
    
        private static final byte[] EXPECTED_EVENT_DATA = "!".getBytes(UTF_8);
        private static final String EXPECTED_EVENT_NAME = "event";
        private static FabricClient fabClient_invoke;
        private static ChannelClient channelClient_invoke;
        static boolean init_judge = true;
    
        public void invokeInit(){
            try {
                Util.cleanUp();
                String caUrl = Config.CA_ORG1_URL;
                CAClient caClient = new CAClient(caUrl, null);
                UserContext adminUserContext = new UserContext();
                adminUserContext.setName(Config.ADMIN);
                adminUserContext.setAffiliation(Config.ORG1);
                adminUserContext.setMspId(Config.ORG1_MSP);
                caClient.setAdminUserContext(adminUserContext);
                adminUserContext = caClient.enrollAdminUser(Config.ADMIN, Config.ADMIN_PASSWORD);
                fabClient_invoke = new FabricClient(adminUserContext);
                channelClient_invoke = fabClient_invoke.createChannelClient(Config.CHANNEL_NAME);
                Channel channel = channelClient_invoke.getChannel();
                Peer peer = fabClient_invoke.getInstance().newPeer("peer2.org1.example.com", "grpc://localhost:7051");
                Orderer orderer = fabClient_invoke.getInstance().newOrderer(Config.ORDERER_NAME, Config.ORDERER_URL);
                channel.addPeer(peer);
                channel.addOrderer(orderer);
                channel.initialize();
                init_judge = false;
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
        public void invoke(String fcn, JSONObject jsonObject){
            try {
                if(init_judge) {
                    invokeInit();
                }
                TransactionProposalRequest request = fabClient_invoke.getInstance().newTransactionProposalRequest();
                ChaincodeID ccid = ChaincodeID.newBuilder().setName(Config.CHAINCODE_1_NAME).build();
                request.setChaincodeID(ccid);
                request.setFcn(fcn);
                String[] arguments = new String[4];
                arguments[0] = jsonObject.getString("identifier");
                arguments[1] = jsonObject.getString("syn_data");
                arguments[2] = jsonObject.getString("id_type");
                arguments[3] = jsonObject.getString("root_id");
                request.setArgs(arguments);
                request.setProposalWaitTime(1000);
                Map<String, byte[]> tm2 = new HashMap<>();
                tm2.put("HyperLedgerFabric", "TransactionProposalRequest:JavaSDK".getBytes(UTF_8));
                tm2.put("method", "TransactionProposalRequest".getBytes(UTF_8));
                tm2.put("result", ":)".getBytes(UTF_8));
                tm2.put(EXPECTED_EVENT_NAME, EXPECTED_EVENT_DATA);
                request.setTransientMap(tm2);
                Collection<ProposalResponse> responses = channelClient_invoke.sendTransactionProposal(request);
                for (ProposalResponse res: responses) {
                    ChaincodeResponse.Status status = res.getStatus();
                    Logger.getLogger(InvokeChaincode.class.getName()).log(Level.INFO,"Invoked initIdentity on "+Config.CHAINCODE_1_NAME + ". Status - " + status);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
> InvokeChaincode3
 
    package org.app.chaincode.invocation;
    
    import org.json.JSONObject;
    import org.app.client.CAClient;
    import org.app.client.ChannelClient;
    import org.app.client.FabricClient;
    import org.app.config.Config;
    import org.app.user.UserContext;
    import org.app.util.Util;
    import org.hyperledger.fabric.sdk.*;
    import java.util.Collection;
    import java.util.HashMap;
    import java.util.Map;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    import static java.nio.charset.StandardCharsets.UTF_8;
    
    public class InvokeChaincode0 {
    
        private static final byte[] EXPECTED_EVENT_DATA = "!".getBytes(UTF_8);
        private static final String EXPECTED_EVENT_NAME = "event";
        private static FabricClient fabClient_invoke;
        private static ChannelClient channelClient_invoke;
        static boolean init_judge = true;
    
        public void invokeInit(){
            try {
                Util.cleanUp();
                String caUrl = Config.CA_ORG1_URL;
                CAClient caClient = new CAClient(caUrl, null);
                UserContext adminUserContext = new UserContext();
                adminUserContext.setName(Config.ADMIN);
                adminUserContext.setAffiliation(Config.ORG1);
                adminUserContext.setMspId(Config.ORG1_MSP);
                caClient.setAdminUserContext(adminUserContext);
                adminUserContext = caClient.enrollAdminUser(Config.ADMIN, Config.ADMIN_PASSWORD);
                fabClient_invoke = new FabricClient(adminUserContext);
                channelClient_invoke = fabClient_invoke.createChannelClient(Config.CHANNEL_NAME);
                Channel channel = channelClient_invoke.getChannel();
                Peer peer = fabClient_invoke.getInstance().newPeer("peer3.org1.example.com", "grpc://localhost:7051");
                Orderer orderer = fabClient_invoke.getInstance().newOrderer(Config.ORDERER_NAME, Config.ORDERER_URL);
                channel.addPeer(peer);
                channel.addOrderer(orderer);
                channel.initialize();
                init_judge = false;
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
        public void invoke(String fcn, JSONObject jsonObject){
            try {
                if(init_judge) {
                    invokeInit();
                }
                TransactionProposalRequest request = fabClient_invoke.getInstance().newTransactionProposalRequest();
                ChaincodeID ccid = ChaincodeID.newBuilder().setName(Config.CHAINCODE_1_NAME).build();
                request.setChaincodeID(ccid);
                request.setFcn(fcn);
                String[] arguments = new String[4];
                arguments[0] = jsonObject.getString("identifier");
                arguments[1] = jsonObject.getString("syn_data");
                arguments[2] = jsonObject.getString("id_type");
                arguments[3] = jsonObject.getString("root_id");
                request.setArgs(arguments);
                request.setProposalWaitTime(1000);
                Map<String, byte[]> tm2 = new HashMap<>();
                tm2.put("HyperLedgerFabric", "TransactionProposalRequest:JavaSDK".getBytes(UTF_8));
                tm2.put("method", "TransactionProposalRequest".getBytes(UTF_8));
                tm2.put("result", ":)".getBytes(UTF_8));
                tm2.put(EXPECTED_EVENT_NAME, EXPECTED_EVENT_DATA);
                request.setTransientMap(tm2);
                Collection<ProposalResponse> responses = channelClient_invoke.sendTransactionProposal(request);
                for (ProposalResponse res: responses) {
                    ChaincodeResponse.Status status = res.getStatus();
                    Logger.getLogger(InvokeChaincode.class.getName()).log(Level.INFO,"Invoked initIdentity on "+Config.CHAINCODE_1_NAME + ". Status - " + status);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
> InvokeChaincode4
 
    package org.app.chaincode.invocation;
    
    import org.json.JSONObject;
    import org.app.client.CAClient;
    import org.app.client.ChannelClient;
    import org.app.client.FabricClient;
    import org.app.config.Config;
    import org.app.user.UserContext;
    import org.app.util.Util;
    import org.hyperledger.fabric.sdk.*;
    import java.util.Collection;
    import java.util.HashMap;
    import java.util.Map;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    import static java.nio.charset.StandardCharsets.UTF_8;
    
    public class InvokeChaincode0 {
    
        private static final byte[] EXPECTED_EVENT_DATA = "!".getBytes(UTF_8);
        private static final String EXPECTED_EVENT_NAME = "event";
        private static FabricClient fabClient_invoke;
        private static ChannelClient channelClient_invoke;
        static boolean init_judge = true;
    
        public void invokeInit(){
            try {
                Util.cleanUp();
                String caUrl = Config.CA_ORG1_URL;
                CAClient caClient = new CAClient(caUrl, null);
                UserContext adminUserContext = new UserContext();
                adminUserContext.setName(Config.ADMIN);
                adminUserContext.setAffiliation(Config.ORG1);
                adminUserContext.setMspId(Config.ORG1_MSP);
                caClient.setAdminUserContext(adminUserContext);
                adminUserContext = caClient.enrollAdminUser(Config.ADMIN, Config.ADMIN_PASSWORD);
                fabClient_invoke = new FabricClient(adminUserContext);
                channelClient_invoke = fabClient_invoke.createChannelClient(Config.CHANNEL_NAME);
                Channel channel = channelClient_invoke.getChannel();
                Peer peer = fabClient_invoke.getInstance().newPeer("peer4.org1.example.com", "grpc://localhost:7051");
                Orderer orderer = fabClient_invoke.getInstance().newOrderer(Config.ORDERER_NAME, Config.ORDERER_URL);
                channel.addPeer(peer);
                channel.addOrderer(orderer);
                channel.initialize();
                init_judge = false;
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
        public void invoke(String fcn, JSONObject jsonObject){
            try {
                if(init_judge) {
                    invokeInit();
                }
                TransactionProposalRequest request = fabClient_invoke.getInstance().newTransactionProposalRequest();
                ChaincodeID ccid = ChaincodeID.newBuilder().setName(Config.CHAINCODE_1_NAME).build();
                request.setChaincodeID(ccid);
                request.setFcn(fcn);
                String[] arguments = new String[4];
                arguments[0] = jsonObject.getString("identifier");
                arguments[1] = jsonObject.getString("syn_data");
                arguments[2] = jsonObject.getString("id_type");
                arguments[3] = jsonObject.getString("root_id");
                request.setArgs(arguments);
                request.setProposalWaitTime(1000);
                Map<String, byte[]> tm2 = new HashMap<>();
                tm2.put("HyperLedgerFabric", "TransactionProposalRequest:JavaSDK".getBytes(UTF_8));
                tm2.put("method", "TransactionProposalRequest".getBytes(UTF_8));
                tm2.put("result", ":)".getBytes(UTF_8));
                tm2.put(EXPECTED_EVENT_NAME, EXPECTED_EVENT_DATA);
                request.setTransientMap(tm2);
                Collection<ProposalResponse> responses = channelClient_invoke.sendTransactionProposal(request);
                for (ProposalResponse res: responses) {
                    ChaincodeResponse.Status status = res.getStatus();
                    Logger.getLogger(InvokeChaincode.class.getName()).log(Level.INFO,"Invoked initIdentity on "+Config.CHAINCODE_1_NAME + ". Status - " + status);
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
    }
### org.app.client
> CAClient
 
    package org.app.client;
    
    import java.lang.reflect.InvocationTargetException;
    import java.net.MalformedURLException;
    import java.util.Properties;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    import org.app.user.UserContext;
    import org.app.util.Util;
    import org.hyperledger.fabric.sdk.Enrollment;
    import org.hyperledger.fabric.sdk.exception.CryptoException;
    import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
    import org.hyperledger.fabric.sdk.security.CryptoSuite;
    import org.hyperledger.fabric_ca.sdk.HFCAClient;
    import org.hyperledger.fabric_ca.sdk.RegistrationRequest;
    
    public class CAClient {
    	String caUrl;
    	Properties caProperties;
    	HFCAClient instance;
    	UserContext adminContext;
    	public UserContext getAdminUserContext() {
    		return adminContext;
    	}
    	public void setAdminUserContext(UserContext userContext) {
    		this.adminContext = userContext;
    	}
    	public CAClient(String caUrl, Properties caProperties) throws MalformedURLException, IllegalAccessException, InstantiationException, ClassNotFoundException, CryptoException, InvalidArgumentException, NoSuchMethodException, InvocationTargetException {
    		this.caUrl = caUrl;
    		this.caProperties = caProperties;
    		init();
    	}
    	public void init() throws MalformedURLException, IllegalAccessException, InstantiationException, ClassNotFoundException, CryptoException, InvalidArgumentException, NoSuchMethodException, InvocationTargetException {
    		CryptoSuite cryptoSuite = CryptoSuite.Factory.getCryptoSuite();
    		instance = HFCAClient.createNewInstance(caUrl, caProperties);
    		instance.setCryptoSuite(cryptoSuite);
    	}
    	public HFCAClient getInstance() {
    		return instance;
    	}
    	public UserContext enrollAdminUser(String username, String password) throws Exception {
    		UserContext userContext = Util.readUserContext(adminContext.getAffiliation(), username);
    		if (userContext != null) {
    			Logger.getLogger(CAClient.class.getName()).log(Level.WARNING, "CA -" + caUrl + " admin is already enrolled.");
    			return userContext;
    		}
    		Enrollment adminEnrollment = instance.enroll(username, password);
    		adminContext.setEnrollment(adminEnrollment);
    		Logger.getLogger(CAClient.class.getName()).log(Level.INFO, "CA -" + caUrl + " Enrolled Admin.");
    		Util.writeUserContext(adminContext);
    		return adminContext;
    	}
    	public String registerUser(String username, String organization) throws Exception {
    		UserContext userContext = Util.readUserContext(adminContext.getAffiliation(), username);
    		if (userContext != null) {
    			Logger.getLogger(CAClient.class.getName()).log(Level.WARNING, "CA -" + caUrl +" User " + username+ " is already registered.");
    			return null;
    		}
    		RegistrationRequest rr = new RegistrationRequest(username, organization);
    		String enrollmentSecret = instance.register(rr, adminContext);
    		Logger.getLogger(CAClient.class.getName()).log(Level.INFO, "CA -" + caUrl + " Registered User - " + username);
    		return enrollmentSecret;
    	}
    	public UserContext enrollUser(UserContext user, String secret) throws Exception {
    		UserContext userContext = Util.readUserContext(adminContext.getAffiliation(), user.getName());
    		if (userContext != null) {
    			Logger.getLogger(CAClient.class.getName()).log(Level.WARNING, "CA -" + caUrl + " User " + user.getName()+" is already enrolled");
    			return userContext;
    		}
    		Enrollment enrollment = instance.enroll(user.getName(), secret);
    		user.setEnrollment(enrollment);
    		Util.writeUserContext(user);
    		Logger.getLogger(CAClient.class.getName()).log(Level.INFO, "CA -" + caUrl +" Enrolled User - " + user.getName());
    		return user;
    	}
    }
> ChannelClient
 
    package org.app.client;
    
    import static java.nio.charset.StandardCharsets.UTF_8;
    import java.io.File;
    import java.io.IOException;
    import java.util.Collection;
    import java.util.HashMap;
    import java.util.Map;
    import java.util.concurrent.CompletableFuture;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    import org.hyperledger.fabric.sdk.BlockEvent.TransactionEvent;
    import org.hyperledger.fabric.sdk.ChaincodeEndorsementPolicy;
    import org.hyperledger.fabric.sdk.ChaincodeID;
    import org.hyperledger.fabric.sdk.Channel;
    import org.hyperledger.fabric.sdk.InstantiateProposalRequest;
    import org.hyperledger.fabric.sdk.Peer;
    import org.hyperledger.fabric.sdk.ProposalResponse;
    import org.hyperledger.fabric.sdk.QueryByChaincodeRequest;
    import org.hyperledger.fabric.sdk.TransactionInfo;
    import org.hyperledger.fabric.sdk.TransactionProposalRequest;
    import org.hyperledger.fabric.sdk.TransactionRequest.Type;
    import org.hyperledger.fabric.sdk.exception.ChaincodeEndorsementPolicyParseException;
    import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
    import org.hyperledger.fabric.sdk.exception.ProposalException;
    
    public class ChannelClient {
    	String name;
    	Channel channel;
    	FabricClient fabClient;
    	public String getName() {
    		return name;
    	}
    	public Channel getChannel() {
    		return channel;
    	}
    	public FabricClient getFabClient() {
    		return fabClient;
    	}
    	public ChannelClient(String name, Channel channel, FabricClient fabClient) {
    		this.name = name;
    		this.channel = channel;
    		this.fabClient = fabClient;
    	}
    	public Collection<ProposalResponse> queryByChainCode(String chaincodeName, String functionName, String[] args)
    			throws InvalidArgumentException, ProposalException {
    		Logger.getLogger(ChannelClient.class.getName()).log(Level.INFO,
    				"Querying " + functionName + " on channel " + channel.getName());
    		QueryByChaincodeRequest request = fabClient.getInstance().newQueryProposalRequest();
    		ChaincodeID ccid = ChaincodeID.newBuilder().setName(chaincodeName).build();
    		request.setChaincodeID(ccid);
    		request.setFcn(functionName);
    		if (args != null)
    			request.setArgs(args);
    		Collection<ProposalResponse> response = channel.queryByChaincode(request);
    		return response;
    	}
    	public Collection<ProposalResponse> sendTransactionProposal(TransactionProposalRequest request)
    			throws ProposalException, InvalidArgumentException {
    		Logger.getLogger(ChannelClient.class.getName()).log(Level.INFO,
    				"Sending transaction proposal on channel " + channel.getName());
    		Collection<ProposalResponse> response = channel.sendTransactionProposal(request, channel.getPeers());
    		for (ProposalResponse pres : response) {
    			String stringResponse = new String(pres.getChaincodeActionResponsePayload());
    			Logger.getLogger(ChannelClient.class.getName()).log(Level.INFO,
    					"Transaction proposal on channel " + channel.getName() + " " + pres.getMessage() + " "
    							+ pres.getStatus() + " with transaction id:" + pres.getTransactionID());
    			Logger.getLogger(ChannelClient.class.getName()).log(Level.INFO,stringResponse);
    		}
    		CompletableFuture<TransactionEvent> cf = channel.sendTransaction(response);
    		Logger.getLogger(ChannelClient.class.getName()).log(Level.INFO,cf.toString());
    		return response;
    	}
    	public Collection<ProposalResponse> instantiateChainCode(String chaincodeName, String version, String chaincodePath,
    			String language, String functionName, String[] functionArgs, String policyPath)
    			throws InvalidArgumentException, ProposalException, ChaincodeEndorsementPolicyParseException, IOException {
    		Logger.getLogger(ChannelClient.class.getName()).log(Level.INFO,
    				"Instantiate proposal request " + chaincodeName + " on channel " + channel.getName()
    						+ " with Fabric client " + fabClient.getInstance().getUserContext().getMspId() + " "
    						+ fabClient.getInstance().getUserContext().getName());
    		InstantiateProposalRequest instantiateProposalRequest = fabClient.getInstance()
    				.newInstantiationProposalRequest();
    		instantiateProposalRequest.setProposalWaitTime(180000);
    		ChaincodeID.Builder chaincodeIDBuilder = ChaincodeID.newBuilder().setName(chaincodeName).setVersion(version)
    				.setPath(chaincodePath);
    		ChaincodeID ccid = chaincodeIDBuilder.build();
    		Logger.getLogger(ChannelClient.class.getName()).log(Level.INFO,
    				"Instantiating Chaincode ID " + chaincodeName + " on channel " + channel.getName());
    		instantiateProposalRequest.setChaincodeID(ccid);
    		if (language.equals(Type.GO_LANG.toString()))
    			instantiateProposalRequest.setChaincodeLanguage(Type.GO_LANG);
    		else
    			instantiateProposalRequest.setChaincodeLanguage(Type.JAVA);
    		instantiateProposalRequest.setFcn(functionName);
    		instantiateProposalRequest.setArgs(functionArgs);
    		Map<String, byte[]> tm = new HashMap<>();
    		tm.put("HyperLedgerFabric", "InstantiateProposalRequest:JavaSDK".getBytes(UTF_8));
    		tm.put("method", "InstantiateProposalRequest".getBytes(UTF_8));
    		instantiateProposalRequest.setTransientMap(tm);
    		if (policyPath != null) {
    			ChaincodeEndorsementPolicy chaincodeEndorsementPolicy = new ChaincodeEndorsementPolicy();
    			chaincodeEndorsementPolicy.fromYamlFile(new File(policyPath));
    			instantiateProposalRequest.setChaincodeEndorsementPolicy(chaincodeEndorsementPolicy);
    		}
    		Collection<ProposalResponse> responses = channel.sendInstantiationProposal(instantiateProposalRequest);
    		CompletableFuture<TransactionEvent> cf = channel.sendTransaction(responses);
    		
    		Logger.getLogger(ChannelClient.class.getName()).log(Level.INFO,
    				"Chaincode " + chaincodeName + " on channel " + channel.getName() + " instantiation " + cf);
    		return responses;
    	}
    	public TransactionInfo queryByTransactionId(String txnId) throws ProposalException, InvalidArgumentException {
    		Logger.getLogger(ChannelClient.class.getName()).log(Level.INFO,
    				"Querying by trasaction id " + txnId + " on channel " + channel.getName());
    		Collection<Peer> peers = channel.getPeers();
    		for (Peer peer : peers) {
    			TransactionInfo info = channel.queryTransactionByID(peer, txnId);
    			return info;
    		}
    		return null;
    	}
    }
> FabricClient
 
    package org.app.client;
    
    import java.io.File;
    import java.io.IOException;
    import java.lang.reflect.InvocationTargetException;
    import java.util.Collection;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    import org.hyperledger.fabric.sdk.ChaincodeID;
    import org.hyperledger.fabric.sdk.Channel;
    import org.hyperledger.fabric.sdk.HFClient;
    import org.hyperledger.fabric.sdk.InstallProposalRequest;
    import org.hyperledger.fabric.sdk.Peer;
    import org.hyperledger.fabric.sdk.ProposalResponse;
    import org.hyperledger.fabric.sdk.User;
    import org.hyperledger.fabric.sdk.exception.CryptoException;
    import org.hyperledger.fabric.sdk.exception.InvalidArgumentException;
    import org.hyperledger.fabric.sdk.exception.ProposalException;
    import org.hyperledger.fabric.sdk.security.CryptoSuite;
    
    public class FabricClient {
    	private HFClient instance;
    	public HFClient getInstance() {
    		return instance;
    	}
    	public FabricClient(User context) throws CryptoException, InvalidArgumentException, IllegalAccessException, InstantiationException, ClassNotFoundException, NoSuchMethodException, InvocationTargetException {
    		CryptoSuite cryptoSuite = CryptoSuite.Factory.getCryptoSuite();
    		// setup the client
    		instance = HFClient.createNewInstance();
    		instance.setCryptoSuite(cryptoSuite);
    		instance.setUserContext(context);
    	}
    	public ChannelClient createChannelClient(String name) throws InvalidArgumentException {
    		Channel channel = instance.newChannel(name);
    		ChannelClient client = new ChannelClient(name, channel, this);
    		return client;
    	}
    	public Collection<ProposalResponse> deployChainCode(String chainCodeName, String chaincodePath, String codepath,
    			String language, String version, Collection<Peer> peers)
    			throws InvalidArgumentException, IOException, ProposalException {
    		InstallProposalRequest request = instance.newInstallProposalRequest();
    		ChaincodeID.Builder chaincodeIDBuilder = ChaincodeID.newBuilder().setName(chainCodeName).setVersion(version)
    				.setPath(chaincodePath);
    		ChaincodeID chaincodeID = chaincodeIDBuilder.build();
    		Logger.getLogger(FabricClient.class.getName()).log(Level.INFO,
    				"Deploying chaincode " + chainCodeName + " using Fabric client " + instance.getUserContext().getMspId()
    						+ " " + instance.getUserContext().getName());
    		request.setChaincodeID(chaincodeID);
    		request.setUserContext(instance.getUserContext());
    		request.setChaincodeSourceLocation(new File(codepath));
    		request.setChaincodeVersion(version);
    		Collection<ProposalResponse> responses = instance.sendInstallProposal(request, peers);
    		return responses;
    	}
    }
### org.app.config
> Config
 
    package org.app.config;
    
    import java.io.File;
    
    public class Config {
    	
    	public static final String ORG1_MSP = "Org1MSP";
    
    	public static final String ORG1 = "org1";
    
    	public static final String ADMIN = "admin";
    
    	public static final String ADMIN_PASSWORD = "adminpw";
    	
    	public static final String CHANNEL_CONFIG_PATH = "config/channel.tx";
    	
    	public static final String ORG1_USR_BASE_PATH = "crypto-config" + File.separator + "peerOrganizations" + File.separator
    			+ "org1.example.com" + File.separator + "users" + File.separator + "Admin@org1.example.com"
    			+ File.separator + "msp";
    	
    	public static final String ORG1_USR_ADMIN_PK = ORG1_USR_BASE_PATH + File.separator + "keystore";
    	
    	public static final String ORG1_USR_ADMIN_CERT = ORG1_USR_BASE_PATH + File.separator + "admincerts";
    	
    	public static final String CA_ORG1_URL = "http://10.108.151.30:7054";
    	
    	public static final String ORDERER_URL = "grpc://10.108.151.30:7050";
    	
    	public static final String ORDERER_NAME = "orderer.example.com";
    	
    	public static final String CHANNEL_NAME = "mychannel";
    	
    	public static final String CHAINCODE_ROOT_DIR = "chaincode";
    	
    	public static final String CHAINCODE_1_NAME = "identity";
    	
    	public static final String CHAINCODE_1_PATH = "github.com/identity";
    	
    	public static final String CHAINCODE_1_VERSION = "1";
    	
    }
### org.app.network
> CreateChannel
 
    package org.app.network;
    
    public class CreateChannel {
        
    }
> DeployInstantiateChaincode
 
    package org.app.network;
    
    public class DeployInstantiateChaincode {
        
    }
> Manager
 
    package org.app.network;
    
    import org.app.client.ChannelClient;
    import org.app.client.FabricClient;
    import org.app.config.Config;
    import org.app.user.UserContext;
    import org.app.util.Util;
    import org.hyperledger.fabric.sdk.*;
    import org.hyperledger.fabric.sdk.security.CryptoSuite;
    import org.json.JSONArray;
    import org.json.JSONException;
    import org.json.JSONObject;
    import java.io.File;
    import java.util.*;
    import java.util.List;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    
    public class Manager {
    
        private UserContext org1Admin;
        private FabricClient fabClient;
        private Orderer orderer;
        private Channel mychannel;
        Peer[] entrys;
        String[] entrys_name;
        String[] entrys_url;
        Peer[] peers;
        String[] peers_name;
        String[] peers_url;
        JSONArray entryArr;
        JSONArray peerArr;
    
        public Manager(JSONObject jsonObject) throws JSONException {
            this.entryArr = jsonObject.getJSONArray("entry");
            this.entrys_name = new String[entryArr.length()];
            this.entrys_url = new String[entryArr.length()];
            for(int i = 0; i < entryArr.length(); i++){
                entrys_name[i] = "entry" + i;
                entrys_url[i] = "grpc://" + entryArr.getJSONObject(i).getString("ip") + ":" + entryArr.getJSONObject(i).getString("port");
            }
            this.peerArr = jsonObject.getJSONArray("peer");
            this.peers_name = new String[peerArr.length()];
            this.peers_url = new String[peerArr.length()];
            for(int i = 0; i < peerArr.length(); i++){
                peers_name[i] = "peer" + i;
                peers_url[i] = "grpc://" + peerArr.getJSONObject(i).getString("ip") + ":" + peerArr.getJSONObject(i).getString("port");
            }
            try {
                CryptoSuite.Factory.getCryptoSuite();
                Util.cleanUp();
                org1Admin = new UserContext();
                File pkFolder1 = new File(Config.ORG1_USR_ADMIN_PK);
                File[] pkFiles1 = pkFolder1.listFiles();
                File certFolder1 = new File(Config.ORG1_USR_ADMIN_CERT);
                File[] certFiles1 = certFolder1.listFiles();
                Enrollment enrollOrg1Admin = Util.getEnrollment(Config.ORG1_USR_ADMIN_PK, pkFiles1[0].getName(),
                        Config.ORG1_USR_ADMIN_CERT, certFiles1[0].getName());
                org1Admin.setEnrollment(enrollOrg1Admin);
                org1Admin.setMspId(Config.ORG1_MSP);
                org1Admin.setName(Config.ADMIN);
                fabClient = new FabricClient(org1Admin);
                orderer = fabClient.getInstance().newOrderer(Config.ORDERER_NAME, Config.ORDERER_URL);
                ChannelConfiguration channelConfiguration = new ChannelConfiguration(new File(Config.CHANNEL_CONFIG_PATH));
                byte[] channelConfigurationSignatures = fabClient.getInstance()
                        .getChannelConfigurationSignature(channelConfiguration, org1Admin);
                mychannel = fabClient.getInstance().newChannel(Config.CHANNEL_NAME, orderer, channelConfiguration,
                        channelConfigurationSignatures);
                Logger.getLogger(CreateChannel.class.getName()).log(Level.INFO, "Channel created - "+ mychannel.getName());
            } catch (Exception e) {
                e.printStackTrace();
            }
            try {
                Logger.getLogger(Util.class.getName()).log(Level.INFO, "Joining - Entrys ");
                entrys = new Peer[entryArr.length()];
                for (int p = 0; p < entrys.length; p++) {
                    entrys[p] = fabClient.getInstance().newPeer(entrys_name[p], entrys_url[p]);
                    mychannel.joinPeer(entrys[p]);
                }
                mychannel.addOrderer(orderer);
                mychannel.initialize();
                Collection peers = mychannel.getPeers();
                Iterator peerIter = peers.iterator();
                while (peerIter.hasNext())
                {
                    Peer pr = (Peer) peerIter.next();
                    Logger.getLogger(CreateChannel.class.getName()).log(Level.INFO,pr.getName()+ " at " + pr.getUrl());
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
            try {
                List<Peer> org1Peers = new ArrayList<>();
                for (int p = 0; p < entrys.length; p++) {
                    org1Peers.add(entrys[p]);
                }
                mychannel.initialize();
                Collection<ProposalResponse> response = fabClient.deployChainCode(Config.CHAINCODE_1_NAME,
                        Config.CHAINCODE_1_PATH, Config.CHAINCODE_ROOT_DIR, TransactionRequest.Type.GO_LANG.toString(),
                        Config.CHAINCODE_1_VERSION, org1Peers);
                int i = 0;
                for (ProposalResponse res : response) {
                    Logger.getLogger(DeployInstantiateChaincode.class.getName()).log(Level.INFO,
                            Config.CHAINCODE_1_NAME + " - Chain code deployment at " + entrys_name[i++]+ " - " + res.getStatus());
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
            try {
                boolean instan_success = false;
                while (instan_success == false){
                    ChannelClient channelClient = new ChannelClient(mychannel.getName(), mychannel, fabClient);
                    String[] arguments = { "" };
                    Collection<ProposalResponse> response = channelClient.instantiateChainCode(Config.CHAINCODE_1_NAME, Config.CHAINCODE_1_VERSION,
                            Config.CHAINCODE_1_PATH, TransactionRequest.Type.GO_LANG.toString(), "init", arguments, null);
                    int i = 0;
                    for (ProposalResponse res : response) {
                        if (res.getStatus() != ChaincodeResponse.Status.SUCCESS) {
                            instan_success = false;
                            break;
                        }
                        instan_success = true;
                        Logger.getLogger(DeployInstantiateChaincode.class.getName()).log(Level.INFO,
                                Config.CHAINCODE_1_NAME + " - Chain code instantiation at " + entrys_name[i++]+ " - " + res.getStatus());
                    }
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
            try {
                Logger.getLogger(Util.class.getName()).log(Level.INFO, "Joining - Peers ");
                peers = new Peer[peerArr.length()];
                for (int p = 0; p < peers.length; p++) {
                    peers[p] = fabClient.getInstance().newPeer(peers_name[p], peers_url[p]);
                    mychannel.joinPeer(peers[p]);
                }
                mychannel.addOrderer(orderer);
                mychannel.initialize();
                Collection peers = mychannel.getPeers();
                Iterator peerIter = peers.iterator();
                while (peerIter.hasNext())
                {
                    Peer pr = (Peer) peerIter.next();
                    Logger.getLogger(CreateChannel.class.getName()).log(Level.INFO,pr.getName()+ " at " + pr.getUrl());
                }
            } catch (Exception e) {
                e.printStackTrace();
            }
        }
        public boolean isOver(){
            return true;
        }
    }
### org.app.user
> CAEnrollment
 
    package org.app.user;
    
    import java.io.Serializable;
    import java.security.PrivateKey;
    import org.hyperledger.fabric.sdk.Enrollment;
    
    public class CAEnrollment implements Enrollment, Serializable {
    	
    	private static final long serialVersionUID = 550416591376968096L;
    	private PrivateKey key;
    	private String cert;
    
    	public CAEnrollment(PrivateKey pkey, String signedPem) {
    		this.key = pkey;
    		this.cert = signedPem;
    	}
    
    	public PrivateKey getKey() {
    		return key;
    	}
    
    	public String getCert() {
    		return cert;
    	}
    }
> RegisterEnrollUser
 
    package org.app.user;
    
    import org.app.client.CAClient;
    import org.app.config.Config;
    import org.app.util.Util;
    
    public class RegisterEnrollUser {
    
    	public static void main(String args[]) {
    		try {
    			Util.cleanUp();
    			String caUrl = Config.CA_ORG1_URL;
    			CAClient caClient = new CAClient(caUrl, null);
    			UserContext adminUserContext = new UserContext();
    			adminUserContext.setName(Config.ADMIN);
    			adminUserContext.setAffiliation(Config.ORG1);
    			adminUserContext.setMspId(Config.ORG1_MSP);
    			caClient.setAdminUserContext(adminUserContext);
    			adminUserContext = caClient.enrollAdminUser(Config.ADMIN, Config.ADMIN_PASSWORD);
    			UserContext userContext = new UserContext();
    			String name = "user"+System.currentTimeMillis();
    			userContext.setName(name);
    			userContext.setAffiliation(Config.ORG1);
    			userContext.setMspId(Config.ORG1_MSP);
    			String eSecret = caClient.registerUser(name, Config.ORG1);
    			userContext = caClient.enrollUser(userContext, eSecret);
    		} catch (Exception e) {
    			e.printStackTrace();
    		}
    	}
    }
> UserContext
 
    package org.app.user;
    
    import java.io.Serializable;
    import java.util.Set;
    import org.hyperledger.fabric.sdk.Enrollment;
    import org.hyperledger.fabric.sdk.User;
    
    public class UserContext implements User, Serializable {
    	
    	private static final long serialVersionUID = 1L;
    	protected String name;
    	protected Set<String> roles;
    	protected String account;
    	protected String affiliation;
    	protected Enrollment enrollment;
    	protected String mspId;
    	
    	public void setName(String name) {
    		this.name = name;
    	}
    
    	public void setRoles(Set<String> roles) {
    		this.roles = roles;
    	}
    
    	public void setAccount(String account) {
    		this.account = account;
    	}
    
    	public void setAffiliation(String affiliation) {
    		this.affiliation = affiliation;
    	}
    
    	public void setEnrollment(Enrollment enrollment) {
    		this.enrollment = enrollment;
    	}
    
    	public void setMspId(String mspId) {
    		this.mspId = mspId;
    	}
    
    	@Override
    	public String getName() {
    		return name;
    	}
    
    	@Override
    	public Set<String> getRoles() {
    		return roles;
    	}
    
    	@Override
    	public String getAccount() {
    		return account;
    	}
    
    	@Override
    	public String getAffiliation() {
    		return affiliation;
    	}
    
    	@Override
    	public Enrollment getEnrollment() {
    		return enrollment;
    	}
    
    	@Override
    	public String getMspId() {
    		return mspId;
    	}
    
    }
### org.app.util
> Util
 
    package org.app.util;
    
    import java.io.BufferedReader;
    import java.io.File;
    import java.io.FileInputStream;
    import java.io.FileOutputStream;
    import java.io.IOException;
    import java.io.InputStream;
    import java.io.InputStreamReader;
    import java.io.ObjectInputStream;
    import java.io.ObjectOutputStream;
    import java.nio.file.Files;
    import java.nio.file.Paths;
    import java.security.KeyFactory;
    import java.security.NoSuchAlgorithmException;
    import java.security.PrivateKey;
    import java.security.spec.InvalidKeySpecException;
    import java.security.spec.PKCS8EncodedKeySpec;
    import java.util.logging.Level;
    import java.util.logging.Logger;
    import javax.xml.bind.DatatypeConverter;
    import org.app.user.CAEnrollment;
    import org.app.user.UserContext;
    import org.hyperledger.fabric.sdk.exception.CryptoException;
    
    public class Util {
    	
    	public static void writeUserContext(UserContext userContext) throws Exception {
    		String directoryPath = "users/" + userContext.getAffiliation();
    		String filePath = directoryPath + "/" + userContext.getName() + ".ser";
    		File directory = new File(directoryPath);
    		if (!directory.exists())
    			directory.mkdirs();
    		FileOutputStream file = new FileOutputStream(filePath);
    		ObjectOutputStream out = new ObjectOutputStream(file);
    		out.writeObject(userContext);
    		out.close();
    		file.close();
    	}
    
    	public static UserContext readUserContext(String affiliation, String username) throws Exception {
    		String filePath = "users/" + affiliation + "/" + username + ".ser";
    		File file = new File(filePath);
    		if (file.exists()) {
    			FileInputStream fileStream = new FileInputStream(filePath);
    			ObjectInputStream in = new ObjectInputStream(fileStream);
    			UserContext uContext = (UserContext) in.readObject();
    			in.close();
    			fileStream.close();
    			return uContext;
    		}
    		return null;
    	}
    	
    	public static CAEnrollment getEnrollment(String keyFolderPath,  String keyFileName,  String certFolderPath, String certFileName)
    			throws IOException, NoSuchAlgorithmException, InvalidKeySpecException, CryptoException {
    		PrivateKey key = null;
    		String certificate = null;
    		InputStream isKey = null;
    		BufferedReader brKey = null;
    		try {
    			isKey = new FileInputStream(keyFolderPath + File.separator + keyFileName);
    			brKey = new BufferedReader(new InputStreamReader(isKey));
    			StringBuilder keyBuilder = new StringBuilder();
    			for (String line = brKey.readLine(); line != null; line = brKey.readLine()) {
    				if (line.indexOf("PRIVATE") == -1) {
    					keyBuilder.append(line);
    				}
    			}
    			certificate = new String(Files.readAllBytes(Paths.get(certFolderPath, certFileName)));
    			byte[] encoded = DatatypeConverter.parseBase64Binary(keyBuilder.toString());
    			PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(encoded);
    			KeyFactory kf = KeyFactory.getInstance("ECDSA");
    			key = kf.generatePrivate(keySpec);
    		} finally {
    			isKey.close();
    			brKey.close();
    		}
    		CAEnrollment enrollment = new CAEnrollment(key, certificate);
    		return enrollment;
    	}
    	
    	public static void cleanUp() {
    		String directoryPath = "users";
    		File directory = new File(directoryPath);
    		deleteDirectory(directory);
    	}
    	
    	  public static boolean deleteDirectory(File dir) {
    	        if (dir.isDirectory()) {
    	            File[] children = dir.listFiles();
    	            for (int i = 0; i < children.length; i++) {
    	                boolean success = deleteDirectory(children[i]);
    	                if (!success) {
    	                    return false;
    	                }
    	            }
    	        }
    	        Logger.getLogger(Util.class.getName()).log(Level.INFO, "Deleting - " + dir.getName());
    	        return dir.delete();
    	    }
    
    }

