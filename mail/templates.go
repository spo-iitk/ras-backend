package mail

const (
	DefaultTemplate = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
	<html dir="ltr" lang="en">
	  <head></head>
	  <body style="background-color:#ffffff;margin:0 auto;font-family:-apple-system, BlinkMacSystemFont, 'Segoe UI', 'Roboto', 'Oxygen', 'Ubuntu', 'Cantarell', 'Fira Sans', 'Droid Sans', 'Helvetica Neue', sans-serif">
		<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation" style="width:100%;margin:0 auto;padding:0px 20px">
		  <tbody>
			<tr style="width:100%">
			  <td>
				<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation" style="margin-top:10px;width:100%">
				  <tbody>
					<tr>
					  <td><img alt="Recruitment Automation System &lt;br /&gt; Indian Institute of Technology, Kanpur" height="75" src="https://raw.githubusercontent.com/spo-iitk/ras-backend/main/.github/images/logo.png" style="display:block;outline:none;border:none;text-decoration:none" width="475" /></td>
					</tr>
				  </tbody>
				</table>
				<h1 style="color:#1d1c1d;font-size:24px;font-weight:700;margin:30px 0;padding:0;line-height:28px">{{.Subject}}</h1>
				<p style="font-size:16px;line-height:20px;margin:16px 0;margin-bottom:30px;white-space:pre-wrap">{{.Body}}</p>
				<p style="font-size:12px;line-height:16px;margin:16px 0;color:#000">This is an auto-generated email. Please do not reply.</p>
				<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation">
				  <tbody>
					<tr>
					  <td>
						<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation" style="margin-bottom:32px;padding-left:8px;padding-right:8px;width:100%">
						  <tbody style="width:100%">
							<tr style="width:100%">
							  <td data-id="__react-email-column" style="width:66%"><img alt="IIT Kanpur" height="48" src="https://upload.wikimedia.org/wikipedia/en/thumb/a/a3/IIT_Kanpur_Logo.svg/1200px-IIT_Kanpur_Logo.svg.png" style="display:block;outline:none;border:none;text-decoration:none" width="48" /></td>
							  <td data-id="__react-email-column">
								<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation">
								  <tbody>
									<tr>
									  <td>
										<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation">
										  <tbody style="width:100%">
											<tr style="width:100%">
											  <td data-id="__react-email-column"><a href="https://twitter.com/IITKanpur" style="color:#067df7;text-decoration:none" target="_blank"><img alt="Twitter" height="32" src="https://th.bing.com/th/id/OIP.YGJYM4pqXxVMHzPYfdLumgHaHa?rs=1&amp;pid=ImgDetMain" style="display:inline;outline:none;border:none;text-decoration:none;margin-left:16px" width="32" /></a></td>
											  <td data-id="__react-email-column"><a href="https://www.facebook.com/spo.iitkanpur/" style="color:#067df7;text-decoration:none" target="_blank"><img alt="Facebook" height="32" src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b8/2021_Facebook_icon.svg/2048px-2021_Facebook_icon.svg.png" style="display:inline;outline:none;border:none;text-decoration:none;margin-left:16px" width="32" /></a></td>
											  <td data-id="__react-email-column"><a href="https://www.linkedin.com/company/students-placement-office-iit-kanpur/" style="color:#067df7;text-decoration:none" target="_blank"><img alt="LinkedIn" height="32" src="https://static-00.iconduck.com/assets.00/linkedin-icon-1024x1024-net2o24e.png" style="display:inline;outline:none;border:none;text-decoration:none;margin-left:16px" width="32" /></a></td>
											</tr>
										  </tbody>
										</table>
									  </td>
									</tr>
								  </tbody>
								</table>
							  </td>
							</tr>
						  </tbody>
						</table>
					  </td>
					</tr>
				  </tbody>
				</table>
				<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation">
				  <tbody>
					<tr>
					  <td><a href="https://spo.iitk.ac.in/" rel="noopener noreferrer" style="color:#b7b7b7;text-decoration:underline;font-size:14px;" target="_blank">Website</a>   |   <a href="https://placement.iitk.ac.in/" rel="noopener noreferrer" style="color:#b7b7b7;text-decoration:underline;font-size:14px;" target="_blank">RAS Portal</a>   |   <a href="https://phdplacement.iitk.ac.in/" rel="noopener noreferrer" style="color:#b7b7b7;text-decoration:underline;font-size:14px;" target="_blank">PhD Portal</a>
						<p style="font-size:12px;line-height:15px;margin:16px 0;color:#b7b7b7;text-align:left;margin-bottom:50px">©2024 Recruitment Automation System. <br />Students&#x27; Placement Office, IIT Kanpur <br /><br />All rights reserved.</p>
					  </td>
					</tr>
				  </tbody>
				</table>
			  </td>
			</tr>
		  </tbody>
		</table>
	  </body>
	</html>`

	// OTPTemplate and RecruitedTemplate are incomplete.

	OTPTemplate = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
<html dir="ltr" lang="en">
<!-- OTP email template content -->
</html>`

	RecruitedTemplate = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
	<meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
	<html dir="ltr" lang="en">
	
	  <head></head>
	
	  <body style="background-color:#ffffff;margin:0 auto;font-family:-apple-system, BlinkMacSystemFont, &#x27;Segoe UI&#x27;, &#x27;Roboto&#x27;, &#x27;Oxygen&#x27;, &#x27;Ubuntu&#x27;, &#x27;Cantarell&#x27;, &#x27;Fira Sans&#x27;, &#x27;Droid Sans&#x27;, &#x27;Helvetica Neue&#x27;, sans-serif">
		<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation" style="max-width:37.5em;margin:0 auto;padding:0px 20px">
		  <tbody>
			<tr style="width:100%">
			  <td>
				<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation" style="margin-top:10px">
				  <tbody>
					<tr>
					  <td><img alt="Recruitment Automation System &lt;br /&gt; Indian Institute of Technology, Kanpur" height="75" src="https://i.ibb.co/YDvnbf5/Screenshot-2024-02-13-000009.png" style="display:block;outline:none;border:none;text-decoration:none" width="475" /></td>
					</tr>
				  </tbody>
				</table>
				<h1 style="color:#1d1c1d;font-size:36px;font-weight:700;margin:30px 0;padding:0;line-height:42px">Congratualtions!</h1>
				<p style="font-size:24px;line-height:28px;margin:16px 0;margin-bottom:30px">You have been recruited for the position of Software Developer at SPO IITK.<br /></p><img alt="IIT Kanpur" height="300" src="https://i.pinimg.com/originals/ee/cc/42/eecc42c92afa81900f655d4328d790c1.gif" style="display:block;outline:none;border:none;text-decoration:none" width="600" />
				<p style="font-size:18px;line-height:24px;margin:16px 0;color:#000">This is an auto-generated email. Please do not reply.</p>
				<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation">
				  <tbody>
					<tr>
					  <td>
						<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation" style="margin-bottom:32px;padding-left:8px;padding-right:8px;width:100%">
						  <tbody style="width:100%">
							<tr style="width:100%">
							  <td data-id="__react-email-column" style="width:66%"><img alt="IIT Kanpur" height="48" src="https://upload.wikimedia.org/wikipedia/en/thumb/a/a3/IIT_Kanpur_Logo.svg/1200px-IIT_Kanpur_Logo.svg.png" style="display:block;outline:none;border:none;text-decoration:none" width="48" /></td>
							  <td data-id="__react-email-column">
								<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation">
								  <tbody>
									<tr>
									  <td>
										<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation">
										  <tbody style="width:100%">
											<tr style="width:100%">
											  <td data-id="__react-email-column"><a href="https://twitter.com/IITKanpur" style="color:#067df7;text-decoration:none" target="_blank"><img alt="Twitter" height="32" src="https://th.bing.com/th/id/OIP.YGJYM4pqXxVMHzPYfdLumgHaHa?rs=1&amp;pid=ImgDetMain" style="display:inline;outline:none;border:none;text-decoration:none;margin-left:32px" width="32" /></a></td>
											  <td data-id="__react-email-column"><a href="https://www.facebook.com/spo.iitkanpur/" style="color:#067df7;text-decoration:none" target="_blank"><img alt="Facebook" height="32" src="https://upload.wikimedia.org/wikipedia/commons/thumb/b/b8/2021_Facebook_icon.svg/2048px-2021_Facebook_icon.svg.png" style="display:inline;outline:none;border:none;text-decoration:none;margin-left:32px" width="32" /></a></td>
											  <td data-id="__react-email-column"><a href="https://www.linkedin.com/company/students-placement-office-iit-kanpur/" style="color:#067df7;text-decoration:none" target="_blank"><img alt="LinkedIn" height="32" src="https://static-00.iconduck.com/assets.00/linkedin-icon-1024x1024-net2o24e.png" style="display:inline;outline:none;border:none;text-decoration:none;margin-left:32px" width="32" /></a></td>
											</tr>
										  </tbody>
										</table>
									  </td>
									</tr>
								  </tbody>
								</table>
							  </td>
							</tr>
						  </tbody>
						</table>
					  </td>
					</tr>
				  </tbody>
				</table>
				<table align="center" width="100%" border="0" cellPadding="0" cellSpacing="0" role="presentation">
				  <tbody>
					<tr>
					  <td><a href="https://spo.iitk.ac.in/" rel="noopener noreferrer" style="color:#b7b7b7;text-decoration:underline" target="_blank">Website</a>   |   <a href="https://placement.iitk.ac.in/" rel="noopener noreferrer" style="color:#b7b7b7;text-decoration:underline" target="_blank">RAS Portal</a>   |   <a href="https://phdplacement.iitk.ac.in/" rel="noopener noreferrer" style="color:#b7b7b7;text-decoration:underline" target="_blank">PhD Portal</a>
						<p style="font-size:12px;line-height:15px;margin:16px 0;color:#b7b7b7;text-align:left;margin-bottom:50px">©2024 Recruitment Automation System. <br />Students&#x27; Placement Office, IIT Kanpur <br /><br />All rights reserved.</p>
					  </td>
					</tr>
				  </tbody>
				</table>
			  </td>
			</tr>
		  </tbody>
		</table>
	  </body>
	</html>`
)