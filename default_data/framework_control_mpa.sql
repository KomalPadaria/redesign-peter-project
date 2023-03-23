
INSERT INTO public.framework_controls (
        "framework_control_uuid",
        "frameworks_uuid",
        "topic",
        "name",
        "domain",
        "best_practices",
        "solution",
        "groups",
        "created_at",
        "updated_at"
    )
VALUES

      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Policies & Procedures', 'OR-1', 'Organizational Security', 'Establish, regularly review, and update upon key changes, the Information Security Management System (ISMS), which is approved by leadership of the organization, which includes the following:
• Control framework
• Governance, Risk and Compliance (GRC)', 'Recommend implementing the following: 
• Reference established Information and Content Security frameworks e.g. MPA Best Practices, ISO27001’s, NIST 800-53, SANS, CoBIT, CSA, CIS, etc.
• Establish an independent team for Information Security, including a governance committee, to develop policies addressing threats, incidents, risks, etc.
• Prepare organization charts and job descriptions to facilitate the designation of roles and responsibilities as it pertains to security', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Risk Management', 'OR-2', 'Organizational Security', 'Establish a formal, documented security risk management program, to include the following:
• Address workflows, assets, and operations
• Apply principles of Confidentiality, Integrity, and Availability (CIA)
• Regularly review and upon key changes
• Conduct a risk assessment annually
• Document decisions on risk management, to include monitoring and reporting remediation status with relevant stakeholders', 'Recommend implementing the following: 
• Define a clear scope for the security risk assessment and modify as necessary
• Incorporate a systematic approach that uses likelihood of risk occurrence, impact to business objectives/content protection and asset classification for assigning priority (e.g. Business Impact Assessment (BIA))
• Risks identified should tie into the business continuity and disaster recovery plans
• Include risks to cloud environments and infrastructure if applicable
• Conduct meetings with management and key stakeholders regularly to identify and document risks
• A formal exception policy
• Document and maintain a Threat Modeling and Analysis process as applicable
• Ensure WFH/remote access content workflow risks are also documented and addressed as applicable
• Leverage NISTIR 8286, FAIR frameworks, or  ISO 3100:2018
• See NIST''s Secure Software Development Framework (SSDF) NIST 800-218 (https://csrc.nist.gov/Projects/ssdf) as an example for Threat Modeling and on how to develop a Secure Software Development Lifecyle (SSDLC) process for coverage of training, requirements, design, development, testing, release and response.', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Personnel Security', 'OR-3', 'Organizational Security', 'Establish and regularly review a policy and process for background screening on all relevant employees, WFH/remote workers, temporary workers, interns and third party workers (e.g. contractors, freelancers, temp agencies etc.), to include the following:
• Perform in accordance with relevant laws, regulations, union bylaws, and cultural considerations
• Retain all signed agreements and results', 'Recommend implementing the following: 
• Use an accredited background screening company', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Personnel Security', 'OR-4', 'Organizational Security', 'Establish and regularly review a process for on-boarding/off-boarding of employees, WFH/remote workers, temporary workers, interns and third party workers (e.g. contractors, freelancers, temp agencies) by performing the following: 

For On-boarding: 
• Perform background screening
• Communicate and require sign-off from all company personnel for all current policies, procedures, and/or client requirements
• Provision physical/digital access as needed
• Complete required training
• Confidentiality agreements, Non-disclosure agreements (NDAs), etc. specifically applied for on-boarding
• Retain all signed agreements 

For Off-boarding:
• Transfer ownership of data & access as required
• De-provision physical/digital access as needed
• Return all company assets/equipment (e.g. keys, fobs, badges, devices, etc.)
• Confidentiality agreements, Non-disclosure agreements (NDAs), etc. specifically applied for off-boarding
• Retain all signed agreements', 'Recommend implementing the following: 
• Apply on a per-project basis as applicable
• For WFH/remote workers, confidentiality agreements are also recommended for other members at the remote location (e.g. roommate, spouse, etc.), where local laws allow
• Review for role/job changes, geographical relocations, and leave of absence as applicable
• Review disciplinary policy as applicable', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Personnel Security', 'OR-5', 'Organizational Security', 'Establish and regularly review a training and awareness program about security policies and procedures and train employees, WFH/remote workers, temporary workers, interns and third party workers (e.g. contractors, freelancers, temp agencies) upon hire and annually, to include the following:
• For executive management and owners, tailor specific training
• Develop tailored training based on job responsibilities (e.g. interaction with content)
• Maintain a log of all training and attendees', 'Recommend implementing the following: 
• Include training for social engineering, ransomware, malware, phishing, WFH/remote working risks, etc.
• Develop a program to test effectiveness of training e.g. phishing campaigns, tabletop exercises, etc.', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Policies & Procedures', 'OR-6', 'Organizational Security', 'Establish and regularly review an Acceptable Use Policy (AUP) governing the use of Internet (e.g. social media and communication activities) and mobile devices (e.g. phones, tablets, laptops, etc.), to include the following:
• Do not share on any social media platform, forum, blog post, or website: information related to pre-release content and related project activities, unless expressed written consent from the client is obtained', 'Recommend implementing the following:
• Use dedicated accounts for marketing purposes', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Personnel Security', 'OR-7', 'Organizational Security', 'Establish and regularly review a policy and procedure to secure content accessed, processed and/or stored at remote sites and locations (i.e. Work From Home (WFH)/remote workers), to include the following:
• Enable MFA for remote access
• WFH/remote workers must be trained on the Remote and Home Working Policy (WFH) and Procedures, as part of their security awareness training to include acknowledgement of Policies and Procedures.
• Define where WFH/remote work is permitted, and where it is not (e.g. home ok, coffee shop not ok)
• The method of remote access to the organization’s internal systems to perform post-production and/or content creation work
• Establish minimum requirements for physical protection of company assets at the remote location
• The use of studio approved pixel streaming remote access (such as, PCoIP, RGS, Parsec, NICE DCV, etc.) that restricts processing and content storage on local endpoint devices.', 'Recommend implementing the following: 
• Restricting unauthorized access to content from others at the remote working location (e.g. roommate, spouse, etc.).
• Requirements and restrictions for the configuration of wireless network services (Note: wired connection is preferred)
• Where feasible, encourage the use of corporate owned devices when content is stored locally on the endpoint device', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Personnel Security', 'OR-8', 'Organizational Security', 'Ensure contracts and/or Service Level Agreements (SLAs) with third-parties and vendors include the following:  
• Disaster Recovery (DR) and Business Continuity Plans (BCP)
• Data handover and disposal upon service termination 
• Risk Management Process
• Background screening
• Confidentiality agreements/NDAs
• Notification if services are outsourced or subcontracted
• Handling and reporting of incidents
• Compliance with applicable data privacy laws
• Cloud deployments', 'Recommend implementing the following: 
• An independent third-party review/audit of the effectiveness of the vendor security and privacy controls is performed (e.g. MPA Best Practices, CSA Star, ISO, SOC 2 Type 2, etc.), to cover the following: Organizational, Operational, Physical, and Technical Security', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Incident Management', 'OR-9', 'Organizational Security', 'Establish and regularly review a formal incident response process, which covers both IT and content incidents/ events, to include the following: 
• Detection
• Notification/ Escalation
• Response
• Evidence/ Forensics
• Analysis
• Remediation
• Reporting and Metrics', 'Recommend implementing the following: 
• Establish a dedicated incident response team
• Apply to cloud deployments (e.g. IaaS, PaaS, SaaS) 
• Apply to employees, WFH/remote workers, temporary workers, interns, third party workers (e.g. contractors, freelancers, temp agencies etc.), and visitors
• Maintain key contact information, including clients
• Notification of affected business partners and clients
• Notification of law enforcement where applicable
• Anonymous reporting where possible
• A corrective action process, to include root cause, lessons learned, preventative measures taken, etc.', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Policies & Procedures', 'OR-10', 'Organizational Security', 'Establish and regularly review formal plans for Business Continuity and Disaster Recovery to include the following:
• Teams responsible for developing and maintaining the Business Continuity (BCP) and Disaster Recovery (DR) Plans.
• Define threats to critical assets, locations, infrastructure, and business operations  (e.g. loss of power or communications, systems failure, natural disasters, pandemics, breach, etc.).
• Notification to affected business partners and clients as applicable.
• Cover Work From Home (WFH)/remote workers, and business functions that are occurring remotely as applicable.', 'Recommend implementing the following for both BCP and DR: 
• Testing procedures of business continuity and disaster recovery processes regularly, to include tabletop exercises if possible
• Base on Recovery Time Objective (RTO) and Recovery Point Objective (RPO)
• Address in Shared Security Responsibility Model (SSRM) if applicable

For Business Continuity, the following is recommended:
• Workarounds, alternate solutions, etc.

For Disaster Recovery, the following is recommended:
• Priorities for recovery procedures, including steps to restore systems
• Cyber security insurance to help mitigate risks from a cyberattack

For template examples refer to SANS: https://www.sans.org/information-security-policy/ and FEMA: https://www.fema.gov/', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Logistics', 'OP-1', 'Operational Security', 'Establish and regularly review a process to receive client assets, to include the following:
• Maintain a receiving log to be filled out by designated personnel upon receipt of deliveries.', 'Recommend implementing the following: 
• For receiving log, include the following information: Name and signature of courier/delivering entity, name and signature of recipient, time and date of receipt
• For assets that can''t be delivered immediately, store in a secure area (e.g. vault, safe, high-security cage, etc.), including overnight deliveries', '{Site}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Logistics', 'OP-2', 'Operational Security', 'Establish and regularly review a process to package assets according to client specifications and destination laws.', 'Recommend implementing the following: 
• Monitor the on-site packaging and loading of content
• Secure containers depending on asset value (e.g. Pelican case with a combination lock)
• Tamper-evident tape, packaging, and/or seals', '{Site}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Logistics', 'OP-3', 'Operational Security', 'Establish a process for shipping client assets, to include the following:
• Maintain a shipping log that includes: time of shipment, recipient name, address of destination, tracking number from shipper.
• Retain shipping log for one year at a minimum.', 'Recommend implementing the following: 
• Generate a work/shipping order to authorize client asset shipments in/out of the facility
• Content awaiting shipment should be in a secure area under camera surveillance', '{Site}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Logistics', 'OP-4', 'Operational Security', 'Establish a process for transport vehicles handling content, to include the following:
• Lock the vehicle at all times
• Ensure packages are out of view', 'Recommend implementing the following: 
• Theft insurance when transporting sensitive assets or as requested by client', '{Site}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Policies & Procedures', 'OP-5', 'Operational Security', 'Establish and regularly review a process and policy for the classification, protection, and handling of data and assets throughout its lifecycle, according to applicable laws and regulations.', 'Recommend implementing the following: 
• Data retention periods
• Classify according to data sensitivity 
• Third-party/ Vendor data sharing responsibilities (e.g. via contract clauses and SSRM)', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Asset Management', 'OP-6', 'Operational Security', 'Establish and regularly review a process for tracking client assets, to include the following:
• Leverage a content asset management system 
• Utilize a unique asset identifier (e.g., barcode, unique ID) in the system, to include the location, time, and date of each asset transaction
• Retain transaction logs for at least one year', 'Recommend implementing the following: 
• Review transaction logs regularly for anomalies
• Implement watermarking as instructed by client (e.g. spoiling, invisible/visible, forensic, etc.)', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Asset Management', 'OP-7', 'Operational Security', 'Establish and regularly review a process to support the handling of client classified high security titles (e.g. Tier 0), to include the following:
• Aliases (e.g., AKA, working title, code name, etc.).
• Access limited to only authorized personnel.', 'Recommend implementing the following:
• Use studio assigned film security title aliases on assets and in asset tracking systems, including lifecycle management (e.g. handling of alias pre vs post release)
• Segregate communications/assets to not include alias and client title
• Individual NDAs/confidentiality agreements as applicable', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Asset Management', 'OP-8', 'Operational Security', 'Establish and regularly review a process for blank media/raw stock to include: 
• Segregation of duties (e.g. between requestor and personnel authorizing check-out, inventory counter and vault staff, etc.)
• Allow access to storage areas (e.g. locked cabinet, safe) to only authorized personnel
• Tagging (e.g. barcode, assign unique identifier) per unit received
• Designating a secure storage area (e.g. locked cabinet, safe)
• Check in/out process to include logging and monitoring', 'Recommend implementing the following:
• Reconciliation on a regular basis (e.g. inventory counts)', '{Site}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Asset Management', 'OP-9', 'Operational Security', 'Establish and regularly review a process to dispose of stock/client assets (e.g. discs, storyboards, scripts, hard drives, etc.) to include:
• Segregation of duties between asset handler/creator and personnel performing the destruction of assets if possible
• Store assets in a secure location/container prior to disposal
• Erasing, degaussing, shredding, or physically destroying before disposal', 'Recommend implementing the following: 
• Destruction be performed onsite
• Destruction be supervised by company personnel, including a sign-off
• When using a third-party company for destruction, obtain a Certificate of Destruction (CoD)
• Complete destruction within 30 days
• Shred bins be locked with openings small enough that a hand cannot fit inside
• Restrict keys to shred bins to authorized personnel only
• Maintain a log of asset disposal for at least one year
• Reference U.S. Department of Defense 5220.22-M for digital shredding and wiping standards', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Access Control', 'PS-1', 'Physical Security', 'Establish and regularly review a process to physically secure all entry/exit points at facilities, to include the following:
• Apply to facility server room, screening room, datacenters, colocations, loading docks, and cloud providers, etc.
• For a datacenter/colocation or cloud provider, proof can be provided via audit reports covering physical security
• Access control segmentation between other businesses and tenants
• Secure and cover windows where content could be visible from the outside
• Apply to WFH/remote locations if applicable', 'Recommend implementing the following:
• Access control segmentation between content areas and other parts of the facility (e.g. administrative offices, waiting rooms, loading docks, courier pickup and drop-off areas, replication, and mastering)
• Attach privacy screens to monitors', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Access Control', 'PS-2', 'Physical Security', 'Establish a process for visitors to include the following:
• Visitor Log
• Retain visitor logs for one year at a minimum, or as local laws allow
• Verification of identity via valid government issued photo ID (e.g. drivers license, passport, etc.)
• NDA/confidentiality agreement for visitors interacting with sensitive content as applicable', 'Recommend implementing the following:
• Visitor log to capture name, company, entry/exit time, reason for visit, person(s) visiting, and signature of visitor
• Visitor badge/sticker
• Conceal the names of previous visitors
• Make visitor badges/stickers easily distinguishable from company personnel badges
• Communicate restrictions of recording/photographing content on premises
• Accompanied by an authorized employee as feasible', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Personnel Security', 'PS-3', 'Physical Security', 'Establish and regularly review and audit the policy and process for replication and distribution facilities, as permitted by local laws, to perform searches of persons, bags, packages, and personal belongings for content/assets at key entry/exit points and as applicable.', 'Recommend also including the following:
• Document any incidents that occur
• Recording/storage devices (e.g. USB thumb drives, digital cameras, cell phones. etc.)
• Use of transparent bags and containers as applicable', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Access Control', 'PS-4', 'Physical Security', 'Establish and regularly review a process to implement Electronic Access Control (EAC) throughout the facility to cover all areas where content is stored, transmitted, or processed, to include the following:
• Designate an individual(s) to authorize facility access
• Assign electronic access to specific facility areas based on job function and responsibilities
• Restrict electronic access system administration to appropriate personnel
• Keep a log that ties the device (e.g. badge, keycard/fob, etc.) to each company personnel
• Store and manage badge, keycard/fob stock securely
• Restrict access to production systems and areas (e.g. vault, server/machine room) to authorized personnel only
• Deploy access control system on a dedicated network separate from production', 'Recommend implementing the following: 
• Set third-party, contractor, etc. to approved timeframe with expiration date (e.g. 90 days)
• Keep records of any changes to access rights', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Access Control', 'PS-5', 'Physical Security', 'Establish and regularly review a process for electronic access logging and monitoring, to include the following:
• Automated alerts for suspicious or unusual events to restricted areas
• Escalation procedures to appropriate personnel
• System enabled logging for all applicable areas
• Retain logs for one year at a minimum, or as local laws allow', 'Recommend implementing the following:
• Review logs regularly for discrepancies', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Monitoring', 'PS-6', 'Physical Security', 'Install and maintain a camera system that captures all facility entry/exit points and restricted areas (e.g. server/machine room, storage areas, vaults, etc.), as local laws allow, to include the following:
• Restrict physical and/or logical access to the surveillance camera console and to camera equipment (e.g. DVRs, NVRs) to authorized personnel only
• Camera positioning and recordings for adequate coverage, image quality, lighting conditions, accurate date and time stamp, and frame rate of surveillance footage
• For a datacenter/colocation or cloud provider, proof can be provided via audit reports 
• Retain footage for at least 90 days, or the maximum time allowed by law, in a secure location', 'Recommend implementing the following: 
• All camera cables and wiring to be discretely hidden from view and not within reach
• Avoid capturing content on display
• Monitor footage during operating hours and immediately investigate detected security incidents
• Test surveillance equipment regularly
• Ensure surveillance equipment functions properly, including an uninterruptable power supply
• All cameras provided by the building are adequate, and footage is accessible
• Apply to WFH/remote worker locations if possible', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Access Control', 'PS-7', 'Physical Security', 'Install and maintain an alarm system that covers all entry/exit points (including emergency exits), windows, loading docks, fire escapes, and restricted areas (e.g. vault, server/machine room, etc.), to include the following:
• Enable the alarm when facility is unsupervised
• Automated alerts
• Escalation configurations and/or procedures to appropriate personnel
• Issue alarm codes and administrator rights to authorized personnel and review users regularly 
• For a datacenter/colocation or cloud provider, proof can be provided via audit reports 
• Test alarm system regularly', 'Recommend implementing the following: 
• Motion sensors to cover sensitive areas (vault, production areas, etc.) 
• Door prop alerts in restricted areas (e.g. vault, server/machine rooms)
• Apply to WFH/remote locations if possible', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Access Control', 'PS-8', 'Physical Security', 'Establish and regularly review a process to manage the distribution of keys to restricted areas to authorized personnel only (e.g. owner, facilities management, etc.), to include the following:
• Implement a check-in/check-out process to track and monitor the distribution of keys
• Maintain a list of company personnel who are allowed to check out keys and review the list regularly
• Regular inventory checks of physical keys and master keys 
• All keys should be stored in a safe location (e.g. lockbox or safe)
• Change the locks when missing keys to restricted areas cannot be accounted for', 'Recommend implementing the following: 
• For a datacenter/colocation or cloud provider, proof can be provided via audit reports', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Monitoring', 'PS-9', 'Physical Security', 'Install and regularly review environmental controls for facilities that contain servers, storage devices, LAN equipment, network communications devices, and storage media to include the following:
• Maintain ideal temperature and humidity settings
• Alerting system for temperatures and humidity levels beyond the set parameters', 'Recommend the following settings: 
• Temperature (Low End): 64.4 F (18 C) 
• Temperature (High End): 80.6 (27 C)
• Moisture (Low End): 40% relative humidity and 41.9 F (5.5 C) dew point
• Moisture (High End): 60% relative humidity and 59 F (15 C) dew point
• For a datacenter/colocation or cloud provider, proof can be provided via audit reports', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-1', 'Technical Security', 'Establish and regularly review a process for data I/O workflows and systems, to include the following:
• Scan all content for viruses and malware prior to ingest onto the network	 
• Dedicated systems for data I/O
• Segmented data I/O network and workflows
• Segregation of duties between data I/O staff and production staff
• Implement separate isolated networks for data I/O and production
• Use dedicated data I/O systems to move content between external networks (Internet) and internal networks (data I/O network, production)
• Content movement must be initiated from the more secure layer: i.e. push/pull content at the data I/O zone to/from Internet; push/pull content at the production network to/from the data I/O zone
• Implement strict (IP and port) layer 2/3 Access Control Lists (ACLs) to allow outbound network requests from the more trusted inner layer, and deny all inbound requests from the less trusted outer layers
• Hardware-encrypted hard drives using Advanced Encryption Standard (AES) 256-bit encryption can also be used to transfer data between production networks and data I/O systems (e.g. ‘air gapped network’)
• Delete content after it has been on the data I/O system for more than 24 hours', 'Recommend implementing the following: 
• Allow listing to restrict content downloads and uploads to only authorized external sources and destinations
• Enable alerts when transfer is complete and/or downloaded
• If Fully Qualified Domain Names (FQDN) are used for allow listing, the firewall should contain a valid Domain Name System (DNS) entry
• WFH/remote workers that ingest content using their machine should always be disconnected from Internet after content download, during production work, and after content upload
• If content is not downloaded or uploaded by WFH/remote workers and is only accessed via a studio approved remote pixel streaming connection (e.g. PCoIP, RGS, Parsec, NICE DCV, etc.), then the previous point is not applicable', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-2', 'Technical Security', 'Place externally accessible servers (e.g. web servers, remote access servers (VPN gateways, remote access brokers, etc.) within a DMZ, VLAN, or a public subnet DMZ within a Virtual Private Cloud (VPC) and not on an internal network, to include the following:
• Isolate virtual or physical servers in the DMZ to provide only one type of service per server (e.g., web server, etc.)
• Implement network controls to restrict access to the internal network from the DMZ, or access from public subnets to private subnets within a VPC (e.g. ACLs, security groups, etc.)
• Maintain an inventory for the external IP addresses and components that are exposed to the Internet', 'Recommend implementing the following: 
• Review network configurations regularly
• Review restrictions regularly (e.g. IP addresses, ACLs, security groups, etc.)', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-3', 'Technical Security', 'Establish and regularly review a process and policy to implement and use dedicated systems for content transfers, to include the following:
• A minimum of AES 256 encryption end-to-end for content at rest and for content in motion.
• Ensure editing stations and content storage servers are not used to directly transfer content
• Disable Virtual Private Network (VPN)/remote access to transfer systems
• Create an approval process to authorize the transfer of content
• Separate content transfer systems from administrative and production networks
• Delete content after it has been on the content transfer devices/systems for more than 24 hours', 'Recommend implementing the following: 
• Use client-approved transfer systems 
• Implement an exception process as needed 
• Send automatic notifications upon outbound content transmission
• Create and maintain a list of users who are responsible for transferring content
• Implement allow listing on content transfer servers to only allow transfers to and from authorized external transfer servers', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-4', 'Technical Security', 'Establish and regularly review a process to secure any point-to-point connection(s) by using dedicated, private connections, and/or encryption, to include the following:
• Connections over the Internet or public networks should be encrypted using site-to-site VPN
• Encrypt communication over private connections (e.g. dark fiber, leased lines, frame relay, MPLS, etc.)
• Use advanced encryption standard (AES 256) or higher for encryption
• Document all point-to-point (e.g. VPN, private fiber, etc.) connections within the organization', 'Recommend implementing the following: 
• Review connections regularly', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Cryptography', 'TS-5', 'Technical Security', 'Establish and regularly review a policy and process to encrypt devices, cloud endpoints, and virtual machines, to include the following:
• Minimum of AES 256 encryption for content at rest and in motion
• File-based encryption: (i.e. encrypting the content)
• Drive-based encryption: (i.e. encrypting the hard drive)
• Send decryption keys, keypad pins, or passwords using an out-of-band communication protocol (i.e., not on the same storage media as the content itself).
• Encryption of backups of sensitive content (AES-256)

For management of encryption keys:
• Access to keys should only be granted to authorized personnel 
• Segregate duties to separate key management from key usage
• All relevant key transactions/activity should be recorded (logged) in the Cryptographic Key Management System (CKMS)
• If applicable, Cloud Service Providers (CSPs) should provide Cloud Service Consumers (CSC) with the ability to manage their own encryption keys', 'Recommend implementing the following:
• For external encrypted drives with keypad pin authentication, enforce self-erase configuration after pre-defined number of invalid attempts

For management of keys, establish procedures for the following activities:
• Generation 
• Distribution 
• Rotation 
• Revocation 
• Destruction
• Deactivation
• Compromise
• Recovery
• Inventory 
• Backup 

For storage, ensure the following: 
• Encrypt encryption key which is at least as strong as the data-encrypting key 
• Store separately from the data-encrypting key
• Store within a secure cryptographic device (e.g. Hardware Security Module (HSM) or a Pin Transaction Security (PTS) point-of-interaction device)', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Cryptography', 'TS-6', 'Technical Security', 'Establish a process for managing Key Delivery Messages (KDMs) and Trusted Devices List (TDL), to include the following:
• Restrict access to the KDM creator and exhibitor only
• Approval and revocation of trusted devices
• Require clients to provide a list of devices that are trusted for content playback and include expiration date
• Only create KDMs for devices on the TDL
• KDM creation and handling be physically and digitally segregated from DCP handling and replication where feasible
• Confirm that devices on the TDL are appropriate based on rights owners’ approval', 'Recommend implementing the following:
• Ensure that encryption key expiration dates conform to client instructions', '{Site}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-7', 'Technical Security', 'Establish and regularly review security baselines, policies, and procedures to configure corporate systems and infrastructure (e.g. laptops, workstations, servers, SAN/NAS, virtual machine infrastructure, WAN, LAN) used at an onsite facility, cloud infrastructure, and for those used by WFH/remote workers, to include the following:
• Install anti-virus/anti-malware
• Disable or remove local accounts on systems or rename username and change the default password
• Disable guest accounts and network shares 
• Remove, uninstall, or disable all unnecessary software and services
• Prohibit users from being administrators on their own workstations, unless required for software
• Block input/output (I/O), mass storage, external storage, and mobile storage devices on all systems that handle or store content, with the exception of systems used for content I/O', 'Recommend implementing the following:
• Enable local firewalls
• Leverage hardening guidelines provided by application providers
• Implement password-protected screensavers or screen-lock software for servers, workstations, cloud endpoints, and WFH/remote workers
• Apply to BYOD where possible', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-8', 'Technical Security', 'Establish and regularly review a process for default administrator accounts and other default accounts, to include the following:
• Identify all default account(s)
• Change the password for all default accounts
• Change the default username(s), when possible', 'Recommend implementing the following: 
• Limit the use of these accounts to special situations that require these credentials (e.g. operating system updates, patch installations, software updates, etc.).
• Apply to WFH/remote workers on equipment, such as firewalls, WIFI, and routers, etc., if possible', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-9', 'Technical Security', 'Establish and regularly review a process for endpoint protection, to include the following:
• Endpoint protection, anti-virus, and anti-malware software with a centralized management console
• Updating anti-virus and anti-malware definitions regularly and performing regular scans on systems
• Apply to WFH/remote worker devices if possible

Apply to the following:
• Workstations (e.g. desktop, laptop)
• Servers
• SAN/NAS
• Virtual Machines
• Cloud infrastructure', 'Recommend implementing the following: 
• Local firewalls where feasible
• Installation of Endpoint Detection and Response (EDR), XDR (Extended Detection and Response), or MXDR (Managed Extended Detection and Response)
• Also apply to Bring Your Own Device (BYOD) where possible', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-10', 'Technical Security', 'Establish and regularly review a process to define security controls and standards for company issued and managed mobile devices (e.g. tablets, cell phones, laptops, etc.), to include the following:
• Report all lost or stolen devices immediately.
• Anti-virus/anti-malware protection 
• Automatic inactivity lock of device during non-use 
• Mobile Device Management (MDM) and/or Mobile Application Management (MAM)
• Ability to conduct a remote wipe should the device be lost, stolen, compromised, etc.
• Require encryption of the entire device', 'Recommend implementing the following:
• Apply to BYOD where possible', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-11', 'Technical Security', 'Implement Security Information and Event Management (SIEM) and regularly review system logs, to include the following:
• Centralized real-time logging of firewalls, authentication servers, network operating systems, content transfer systems, remote access mechanisms, virtual machines/servers, storage services, databases, container-based application services, API gateway connections, key generation/management, etc.
• Retain logs for a period of one year, where local laws permit
• Access to logging infrastructure should be restricted to authorized personnel only
• A synchronized time service protocol (e.g. Network Time Protocol (NTP)) to ensure all systems have a correct and consistent time
• Protect logs from unauthorized deletion or modification by applying appropriate access rights on log files
• Configure logging systems to send automatic notifications when security events are detected.
• Assign personnel to review logs and respond to alerts
• Incorporate into BCP & Incident Response procedures.', 'Recommend implementing the following: 
• Enable local logging on isolated systems
• Include logging and monitoring of spikes in resource utilization and capacity management
• Log, monitor, and review all authentication activity and alerts

Alert configurations should include the following:
• Successful and unsuccessful attempts to connect to the content/production network
• Unusual file size and/or time of day transport of content 
• Repeated attempts for unauthorized file access 
• Attempts at privileged access', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-12', 'Technical Security', 'Document the network and cloud infrastructure and topology diagrams, and update when significant changes are made.', 'Recommend implementing the following:
• Including WAN, DMZ, LAN, WLAN (wireless), VLAN, firewalls, switches, endpoints, remote access, etc.', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-13', 'Technical Security', 'Establish a policy to use layer 3 switches/devices to manage network traffic, to include the following:
• Port security to be enabled
• Disable unused ports on switches
• Disable Simple Network Management Protocol (SNMP) if it is not in use. Use SNMP v3 or higher with strong passwords for community strings', 'Recommend implementing the following:
• Use device administrator credentials with strong passwords
• Use physical ethernet cable locks to ensure that a network cable cannot be connected to an alternate/unauthorized device
• Network-based access control, i.e. 802.1X
• If layer 2 switches are still in use, ensure that a higher layer network communications device is providing network isolation/traffic control 
• Restrict the use of non-switched devices such as hubs and repeaters', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-14', 'Technical Security', 'Establish and regularly review a process and policy to separate external network(s)/WAN(s) from the internal network(s) by using stateful inspection firewall(s), to include the following:
• Review firewall Access Control Lists (ACLs) regularly 
• WFH/remote locations to have a firewall to segregate the WAN (Internet) from the internal network used to access content as applicable

Apply the following configurations:
• Firewalls with Access Control Lists that deny all WAN traffic to any internal network other than to explicit hosts that reside on the DMZ
• Firewall WAN network to prohibit direct network access to the internal content/production network
• Firewall rules to generate logs for all traffic and for all configuration changes, and logs should be inspected regularly
• Deny all incoming and outgoing network requests by default
• Enable only explicitly defined incoming requests by specific protocol and destination
• Enable only explicitly defined outgoing requests by specific protocol and source
• For externally accessible hosts, only allow incoming requests to needed ports 
• Restrict unencrypted communication protocols e.g. Telnet and FTP, and replace with encrypted versions
• Firewall to have a subscription to anti-virus and intrusion detection updates
• Deploy a Web Application Firewall (WAF) in front of Internet facing web applications and APIs', 'Recommend implementing the following: 
• Anti-spoofing filters
• Block the following: non-routable IP addresses internal addresses over external ports, UDP and ICMP echo requests, unused ports and services, and unauthorized DNS zone transfers', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-15', 'Technical Security', 'Establish and regularly review a process to isolate the content/production networks from non-content/production networks (e.g. office network, DMZ, content transfer, Internet etc.), to include the following:
• Layer 1 physical air gap if applicable
• Logical segmentation via Layer 2 or Layer 3 VLAN ACLs
• Prohibit bridging or dual-homed networking (physical network bridging) on computer systems between content/production networks and non-content/production networks.
• If applicable to WFH/remote locations, segregate production network through a remote connection via client approved remote access (e.g. PCoIP, RGS, Parsec, NICE DCV, etc.)', 'Recommend implementing the following: 
• Review network configurations regularly
• Update upon key changes', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-16', 'Technical Security', 'Establish and regularly review a process and policy for firewall management, to include the following:
• Provisioning requirements based off the concept of least privilege
• Change control requirements (e.g. patching, upgrades, firewall rule management)
• Do not allow direct firewall management from the Internet or WAN
• Require secure remote access with MFA for administration 
• Configure to alert key security events', 'Recommend implementing the following: 
• Review role access regularly
• Review alert configuration regularly
• Update upon key changes', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-17', 'Technical Security', 'Establish a policy to implement a network-based intrusion detection/prevention system (IDS/IPS) to protect the network, to include the following:
• Configure the system to alert and block suspicious network activity
• Implement basic border gateway services (e.g. gateway anti-virus, and URL filtering)
• Update attack signature definitions/policies regularly
• Log all activity and configuration changes', 'Recommend implementing the following: 
• Consider host-based intrusion detection systems
• Utilize virtual patching', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-18', 'Technical Security', 'Establish and regularly review a process and policy for Internet access in production networks and all systems that process or store digital content, to include the following:
• Prohibit directly accessing unauthorized Internet sites, resources, or services.
• Prohibit direct email access
• Implement firewall rules to deny all outbound traffic by default, including to the Internet and other internal networks', 'If a business case requires Internet access from the production network, the following is recommended: 
• For cases where services  (e.g. anti-virus definitions, patches, licenses, etc.) are needed on the production network, explicitly allow protocols and ports (i.e. layer 2/3 ACLs) that require connections to the services. 
• If Internet is needed, proxy servers must be used to broker access

For isolated web browsing/email access, the following is recommended: 
• Browser isolation tools via a virtual environment that is not on the production network (e.g. Ericom RBI, McAfee Light Point, Zscaler, Palo Alto Prisma, Menlo Browser Isolation, etc.)

For use of KVM for web browsing and/or email access, the following is recommended:
• A keyboard/video/mouse (KVM) solution to a machine with Internet access not connected to the production network 
• Ensure that any physical ports on the KVM switch which are not in use are properly locked down', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-19', 'Technical Security', 'Establish and regularly review a policy to enforce authorization and authentication policy of employees, WFH/remote workers, temporary workers, interns and third party workers (e.g. contractors, freelancers, temp agencies), administrative accounts, service accounts, to include the following: 
• Unique username
• Use the principles of least privilege

For passwords and passphrases:
• Minimum password or passphrase length of at least 12 characters
• Minimum of 3 of the following parameters: upper case, lower case, numeric, or special characters
• Maximum password or passphrase age of 1 year (not applicable to service accounts)
• Minimum password or passphrase age of 1 day (not applicable to service accounts)
• Maximum of 5 invalid logon attempts
• User accounts locked after invalid logon attempts must be manually unlocked by a system administrator
• Can''t reuse last 5 passwords or passphrases (not applicable to service accounts)
• Changing of password or passphrase upon detection of suspicious activity or incident

For Multi-Factor Authentication (MFA), only apply to the following:
• All administrative accounts
• Any Internet facing systems, including webmail, web portal, cloud portal
• WFH/remote workers when connecting to corporate and/or production systems
• Source code repository', 'Recommend implementing the following: 
• Apply MFA to all accounts where feasible

For administrator and service accounts, the following is recommended:
• Ensure accounts are still used for intended purposes only (e.g. database queries, application-to-application communication, etc.)
• Monitoring and central logging of successful logons, failed logons, and lockouts
• Privileged Account Management (PAM) tool', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-20', 'Technical Security', 'Establish and regularly review a process to manage access to all information systems for employees, WFH/remote workers, temporary workers, interns and third party workers (e.g. contractors, freelancers, temp agencies), administrative accounts, service accounts, to include the following:
• Implement Identity Access Management (IAM) (e.g. role-based access control (RBAC), attribute-based access control (ABAC), single sign on system, identity federation standards, and directory service (e.g. Active Directory, Open Directory, LDAP, Zero Trust Architecture))
• Assign dedicated personnel to manage access', 'Recommend implementing the following: 
• Where applicable, use of cloud hosted directory services (e.g. JumpCloud, OKTA, Azure Active Directory, AWS Directory Service, etc.)
• Configure systems and applications to log administrator actions and record, at the minimum, the following information: username, time stamp, action, additional information (action parameters)
• Regularly review for discrepancies, and unusual or suspicious activity', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-21', 'Technical Security', 'Establish and regularly review a process to enable MFA for remote user access to all environments, to include the following:
• Enable AES 256 encryption for all tiers
• Third-party IT service provider access to be limited to a specific time frame
• Remote access accounts to not be shared (use individual, unique accounts)
• Avoid use of the following methods for remote access: FTP, Telnet
• Remote access to be logged and reviewed real time with alerts generated for suspicious activity 
• VPN configuration to not allow split tunneling

Follow the below tier structure:
• Tier 1: Access only to a corporate network or service that doesn’t store content (e.g. VPN to corporate VLAN for file share access, webmail, Office365, etc.)
• Tier 2: WFH/remote worker access to a content production network via studio approved pixel streaming (e.g. PCoIP, RGS, Parsec, NICE DCV, etc.). Do not allow any access to copy content files to the local machine. Access to a production network is only be granted via an access broker that is on a non-production network (e.g. DMZ)
• Tier 3: Elevated VPN administrative access to a production network for approved personnel to perform their job responsibilities. Use a launchpad/bastion host as an intermediate machine (‘jump box’) from a non-production network, to connect to the production network, without any direct connection to production allowed from the Internet', 'Recommend implementing the following: 
• Maintain a list of authorized remote access users
• Regularly review user list for discrepancies, and unusual or suspicious activity', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-22', 'Technical Security', 'Establish and regularly review a process for corporate web filtering to address the following: 
• Peer-to-peer file sharing
• Malware/ransomware
• Malicious sites', 'Recommend implementing the following: 
• Use of DNS filtering
• Use of a CASB (Cloud Access Security Broker) to monitor and restrict cloud software usage and access
• If applicable, also cover WFH/remote workers and BYOD devices, as local laws permit', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-23', 'Technical Security', 'Establish and regularly review a process for corporate email filtering to detect, report, and block the following: 
• Phishing emails
• Malware/ransomware
• Transmission of sensitive asset/content material
• Executable file attachments', 'Recommend implementing the following: 
• If applicable, also cover WFH/remote workers and BYOD devices, as local laws permit
• Incorporate into Incident Management process for reporting', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-24', 'Technical Security', 'Establish and regularly review a process and policy for managing wireless network configurations in all environments to include the following:
• Disable WEP/WPA
• Enable WPA2-PSK (AES), and/or WPA3- SAE
• Change default administrator logon credentials
• Change default network name Service Set Identifier (SSID) and use non-company, non-production names
• Set a complex wireless access point passphrase and change regularly
• Remote Authentication Dial In User Service (RADIUS) for authentication (does not apply to guest networks)
• Disconnect wireless Network Interface Cards (NICs) from production computers
• Segregate guest networks from the company’s other networks
• Restrict guest networks to access only the Internet', 'Recommend implementing the following: 
• Use WPA2-Enterprise (AES) if applicable
• MAC address filtering and disallow wireless MAC addresses of production devices
• Configure the wireless access point/controller to broadcast only within the required range
• Port-based network access control (e.g. 802.1X framework for wireless networking)
• Lightweight Directory Access Protocol (LDAP) server, such as Active Directory, to manage user accounts
• Public Key Infrastructure to generate and manage client and server certificates
• Configure WPA2 or WPA3 with CCMP (AES)
• Scan for rogue wireless access points and/or use a centralized wireless access to alert rogue connections
• Apply to WFH/remote worker wireless networks and disconnect wireless networks while accessing content locally', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Vulnerability Management', 'TS-25', 'Technical Security', 'Establish and regularly review a process and policy for Vulnerability Management, including vulnerability scans for both internal and external networks, cloud deployments, and virtual machines/containers, to include the following: 
• For external IP ranges and hosts, perform scans monthly at a minimum
• For internal IP ranges and hosts, perform scans quarterly at a minimum
• Investigate and have a remediation plan for issues
• Perform a vulnerability scan after any major application or cloud infrastructure change
• Apply internal scan to WFH/remote worker endpoints where possible 

Also scan the following if applicable:
• Production networks
• Non-Production networks 
• Application Programming Interfaces (APIs)', 'Recommend implementing the following: 
• Investigate and have a remediation plan for critical issues within 48 hours
• Authenticated and unauthenticated scanning
• Leverage Open Web Application Security Project (OWASP)', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Vulnerability Management', 'TS-26', 'Technical Security', 'Establish and regularly review a process and policy to perform penetration testing of all external IP ranges, hosts, web applications, and cloud deployments (if applicable), to include the following:
• Conduct on an annual basis at a minimum
• Investigate and have a remediation plan for issues
• Conducted by an independent third-party or internal red team
• Perform a penetration test after any major application or cloud infrastructure change

Also test the following if applicable:
• Application Programming Interfaces (APIs)', 'Recommend implementing the following: 
• Investigate and have a remediation plan for critical issues within 48 hours
• Authenticated and unauthenticated testing
• Leverage Open Web Application Security Project (OWASP)', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Vulnerability Management', 'TS-27', 'Technical Security', 'Establish and regularly review a process to patch endpoints, servers, applications, virtual machines, network infrastructure devices (e.g. firewalls, routers, switches, etc.), Storage Area Networks (SAN), and Network Attached Storage (NAS), to include the following:
• Investigate and address patches
• Subscribe to security and patch notifications from vendors, other third parties, and security advisories
• Decommission legacy systems that are no longer supported', 'Recommend implementing the following: 
• Investigate and have a remediation plan for critical patches within 48 hours
• A centrally managed patch management system
• Also apply to BYOD where possible', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Policies & Procedures', 'TS-28', 'Technical Security', 'Establish, document and regularly review a process for change control to ensure data, applications, network, and system component updates and changes have been reviewed and approved as required, to include the following:
• Maintain an up-to-date inventory of systems (e.g. Configuration Management Database (CMDB)), system components, and software
• Identify all impacted computer software, data files, database entities, infrastructure, and cloud systems 
• Identify and manage risks, include security controls associated with changes to data, applications, network infrastructure and systems
• Document and retain all changes, test results, and management approvals
• Ensure that appropriate backup or roll-back procedures are documented and tested', 'Recommend implementing the following: 
• Establish a Change Control Board (CCB), consisting of individuals responsible for reviewing and approving any updates or changes', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-29', 'Technical Security', 'Establish and regularly review a process to verify, restrict, and manage access to web and cloud portals, to include the following:
• Use HTTPS signed by a certificate authority (CA)
• Ensure HTTPS certificates are not expired
• For HTTPS, enforce use of a strong cipher suite (e.g. TLS v1.2 or higher)
• Place the web or cloud portal on a dedicated server in the DMZ
• Establish user permissions according to roles (e.g. ability to upload/download content)
• Segregated access between client tenants
• Do not use persistent cookies or cookies that store credentials in plaintext', 'Recommend implementing the following: 
• For sensitive content, set access to expire automatically at predefined intervals, where configurable
• Review user access list to the client web/cloud portal regularly 
• HTTP Strict Transport Security (HSTS)', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-30', 'Technical Security', 'Establish and regularly review a process to provide Shared Security Responsibility Model (SSRM) guidance to the Cloud Service Provider (CSP) and Cloud Service Consumer (CSC).', 'Recommend implementing the following:
• Regularly review the SSRM for updates and changes
• CSC to engage with the CSP to address any issues identified, and SSRM changes to be incorporated into the CSC''s implementation plans
• CSCs to implement the finalized SSRM controls and test the controls to validate the proper operation of CSC security controls 
• For non-conformance issues, develop a plan to remediate', '{Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-31', 'Technical Security', 'Establish and regularly review a process to configure applications and infrastructures, so that Cloud Service Provider (CSP) and Cloud Service Consumer (CSC) user access and intra-tenant access is segregated between tenants (e.g. physically or logically).', 'Recommend implementing the following:
• Compliance with legal, statutory, and regulatory compliance obligations
• Monitor segmentation between intra-tenant access
• For intra-tenant segregation at data center/colocation or cloud provider, proof can be provided via audit reports', '{Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Network Security', 'TS-32', 'Technical Security', 'Establish and regularly review a process to monitor, encrypt, and restrict network connections between environments to only authenticated and authorized connections, to include the following:
• Detect unauthorized connections
• Remove unauthorized connections', 'Recommend implementing the following:
• Regularly review network connections between environments for unauthorized connections', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-33', 'Technical Security', 'Establish and regularly review a process for the detection and correction of cloud misconfigurations, to include the following:
• Proactive alerts
• Appropriate role(s) for reviewing and correcting misconfigurations
• A configuration and management tool
• Investigate and have a remediation plan for misconfigurations', 'Recommend implementing the following:
• Investigate and have a remediation plan for critical misconfigurations within 48 hours
• Utilize scanning tools to detect misconfigurations', '{Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-34', 'Technical Security', 'Establish and regularly review a process for Secure Software Development Lifecycle (SSDLC) for application design, development, deployment, and including a software testing strategy, regardless of location, to include the following:
• Perform a secure code review for each build
• Perform application and code repository security testing
• Include scanning (e.g. TFSEC) coverage for Continuous Integration (CI)/Continuous Delivery (CD) automated pipelines and deployments (e.g. Terraform)
• Include scanning open source libraries when applicable
• Investigate and have a remediation plan for issues', 'Recommend implementing the following:
• Engage a third-party to conduct an independent review of the code if possible
• Document and restrict the results of the secure code review to authorized personnel only
• Perform automated code scans
• If manual scanning tools or code analysis tools are used, a scan should be conducted at each code change and/or production code push', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-35', 'Technical Security', 'Establish and regularly review a process to develop systems and applications based upon principles of Security by Design (SbD) and Privacy by Design (PbD), to include the following:
• Data protection and privacy requirements be included by default at the design stage and throughout the product development lifecycle
• Follow applicable regional/local privacy laws', 'Recommend implementing the following:
• Design documentation to describe how data is protected
• Data Loss Prevention tools (DLP)', '{Site, Cloud}', current_timestamp, current_timestamp),
      (gen_random_uuid(), (select frameworks_uuid from frameworks where name = 'MPA'), 'Information Systems', 'TS-36', 'Technical Security', 'Establish a process for all forms of code, including source code and executable code, to include the following:
• Store in a secure repository
• Access based on the principle of least privilege
• Ensure that credentials and secrets are never embedded in code', 'Recommend implementing the following: 
• A secrets management service to rotate, manage, and retrieve database credentials or secrets
• Credentials and sensitive data be encrypted by a KMS', '{Site, Cloud}', current_timestamp, current_timestamp);
