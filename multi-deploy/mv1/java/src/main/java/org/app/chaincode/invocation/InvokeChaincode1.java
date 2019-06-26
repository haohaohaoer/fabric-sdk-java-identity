package org.app.chaincode.invocation;

import org.app.client.CAClient;
import org.app.client.ChannelClient;
import org.app.client.FabricClient;
import org.app.config.Config;
import org.app.user.UserContext;
import org.app.util.Util;
import org.hyperledger.fabric.sdk.*;
import org.json.JSONObject;

import java.util.Collection;
import java.util.HashMap;
import java.util.Map;
import java.util.logging.Level;
import java.util.logging.Logger;

import static java.nio.charset.StandardCharsets.UTF_8;

public class InvokeChaincode1 {

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
            // Enroll Admin to Org1MSP
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
            //EventHub eventHub = fabClient_invoke.getInstance().newEventHub("eventhub01", "grpc://localhost:7053");
            Orderer orderer = fabClient_invoke.getInstance().newOrderer(Config.ORDERER_NAME, Config.ORDERER_URL);
            channel.addPeer(peer);
            //channel.addEventHub(eventHub);
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
